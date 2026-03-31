// File: internal/web/server.go
package web

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"net/http"
	"strings"
	"sync"

	"keyscome.io/petal/internal/config"
	"keyscome.io/petal/internal/env"
	"keyscome.io/petal/internal/executor"
)

//go:embed static
var staticFiles embed.FS

// Server holds the HTTP server state.
type Server struct {
	taskFile *config.TaskFile
	envFile  string
}

// NewServer creates a new web Server.
func NewServer(tf *config.TaskFile, envFile string) *Server {
	return &Server{taskFile: tf, envFile: envFile}
}

// Start begins listening on addr (e.g. ":8080").
func (s *Server) Start(addr string) error {
	mux := http.NewServeMux()

	// Static files (index.html etc.)
	staticFS, err := fs.Sub(staticFiles, "static")
	if err != nil {
		return fmt.Errorf("embed fs error: %w", err)
	}
	mux.Handle("/", http.FileServer(http.FS(staticFS)))

	// API
	mux.HandleFunc("/api/tasks", s.handleTasks)
	mux.HandleFunc("/api/run/", s.handleRun)

	fmt.Printf("🌱 Petal web UI listening on http://localhost%s\n", addr)
	return http.ListenAndServe(addr, mux)
}

// taskSummary is the JSON shape returned by GET /api/tasks.
type taskSummary struct {
	Name  string   `json:"name"`
	Hosts []string `json:"hosts"`
	Cmd   string   `json:"cmd"`
}

func (s *Server) handleTasks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	summaries := make([]taskSummary, 0, len(s.taskFile.Tasks))
	for _, t := range s.taskFile.Tasks {
		summaries = append(summaries, taskSummary{
			Name:  t.Name,
			Hosts: t.Hosts,
			Cmd:   t.Cmd,
		})
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(summaries)
}

// sseWriter sends Server-Sent Events to the browser.
type sseWriter struct {
	w       http.ResponseWriter
	flusher http.Flusher
	mu      sync.Mutex
}

func newSSEWriter(w http.ResponseWriter) (*sseWriter, bool) {
	f, ok := w.(http.Flusher)
	if !ok {
		return nil, false
	}
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("X-Accel-Buffering", "no")
	return &sseWriter{w: w, flusher: f}, true
}

func (sw *sseWriter) send(msgType, text string) {
	sw.mu.Lock()
	defer sw.mu.Unlock()
	data, _ := json.Marshal(map[string]string{"type": msgType, "text": text})
	fmt.Fprintf(sw.w, "data: %s\n\n", data)
	sw.flusher.Flush()
}

func (sw *sseWriter) done() {
	sw.mu.Lock()
	defer sw.mu.Unlock()
	fmt.Fprintf(sw.w, "data: {\"done\":true}\n\n")
	sw.flusher.Flush()
}

func (s *Server) handleRun(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	taskName := strings.TrimPrefix(r.URL.Path, "/api/run/")
	if taskName == "" {
		http.Error(w, "task name required", http.StatusBadRequest)
		return
	}

	var found *config.RemoteTask
	for i := range s.taskFile.Tasks {
		if s.taskFile.Tasks[i].Name == taskName {
			t := s.taskFile.Tasks[i]
			found = &t
			break
		}
	}
	if found == nil {
		http.Error(w, "task not found", http.StatusNotFound)
		return
	}

	sw, ok := newSSEWriter(w)
	if !ok {
		http.Error(w, "streaming not supported", http.StatusInternalServerError)
		return
	}

	s.runTask(sw, *found)
	sw.done()
}

// runTask executes a RemoteTask and streams output via SSE.
func (s *Server) runTask(sw *sseWriter, t config.RemoteTask) {
	sw.send("header", fmt.Sprintf("=== Task: %s ===", t.Name))

	taskEnv := env.MergeFromStringMap(t.Env)

	fileEnv := map[string]env.EnvVar{}
	if s.envFile != "" {
		loaded, err := env.LoadEnvFile(s.envFile)
		if err != nil {
			sw.send("error", fmt.Sprintf("[Task: %s] Failed to load env file: %v", t.Name, err))
			return
		}
		fileEnv = loaded
	}

	globalEnv := env.MergeFromStringMap(s.taskFile.Env)
	resolver := env.NewResolver(globalEnv, fileEnv, taskEnv)
	envAssign := buildEnvString(resolver.Flat(false))

	renderedCmdRaw := resolver.Render(t.Cmd, false)
	cmdToRun := env.StripShellComments(renderedCmdRaw)
	maskedDisplay := resolver.Render(t.Cmd, true)

	sw.send("header", fmt.Sprintf(">>> Rendered Script on %s:", firstHost(t.Hosts)))
	sw.send("running", maskedDisplay)

	if len(t.Hosts) == 0 {
		sw.send("error", fmt.Sprintf("[Task: %s] No hosts defined", t.Name))
		return
	}

	exec := executor.SSHExecutor{}
	var wg sync.WaitGroup
	type result struct {
		host   string
		output string
		err    error
	}
	results := make(chan result, len(t.Hosts))

	for _, host := range t.Hosts {
		wg.Add(1)
		go func(h string) {
			defer wg.Done()
			out, err := exec.Execute(h, envAssign, cmdToRun)
			results <- result{host: h, output: out, err: err}
		}(host)
	}

	// Close results channel once all goroutines finish.
	go func() {
		wg.Wait()
		close(results)
	}()

	for res := range results {
		if res.err != nil {
			sw.send("error", fmt.Sprintf("[%s][%s] Error", t.Name, res.host))
			sw.send("error", res.output)
		} else {
			sw.send("success", fmt.Sprintf("[%s][%s] Success", t.Name, res.host))
			sw.send("success", res.output)
		}
	}

	sw.send("header", fmt.Sprintf("=== Completed: %s ===", t.Name))
}

func buildEnvString(envMap map[string]string) string {
	if len(envMap) == 0 {
		return ""
	}
	parts := make([]string, 0, len(envMap))
	for k, v := range envMap {
		parts = append(parts, fmt.Sprintf("%s='%s'", k, v))
	}
	return strings.Join(parts, " ")
}

func firstHost(hosts []string) string {
	if len(hosts) == 0 {
		return "(none)"
	}
	return hosts[0]
}
