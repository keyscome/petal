// File: internal/task/runner.go
package task

import (
	"fmt"
	"strings"
	"sync"

	"keyscome.io/petal/internal/color"
	"keyscome.io/petal/internal/config"
	"keyscome.io/petal/internal/env"
)

type Executor interface {
	Execute(host, envAssign, cmd string) (string, error)
}

type Runner struct {
	GlobalEnv map[string]env.EnvVar
	EnvFile   string
	Exec      Executor
}

func NewRunner(globalEnv map[string]env.EnvVar, envFile string, exec Executor) *Runner {
	return &Runner{
		GlobalEnv: globalEnv,
		EnvFile:   envFile,
		Exec:      exec,
	}
}

func (r *Runner) Run(task config.RemoteTask) {
	color.PrintHeader("=== Task: %s ===", task.Name)

	// Convert task.Env
	taskEnv := env.MergeFromStringMap(task.Env)

	// Load file env
	fileEnv := map[string]env.EnvVar{}
	if r.EnvFile != "" {
		loaded, err := env.LoadEnvFile(r.EnvFile)
		if err != nil {
			color.PrintError("[Task: %s] Failed to load env file: %v", task.Name, err)
			return
		}
		fileEnv = loaded
	}

	// Build env resolver
	resolver := env.NewResolver(r.GlobalEnv, fileEnv, taskEnv)
	envAssign := buildEnvString(resolver.Flat())
	renderedCmd := resolver.Render(task.Cmd)

	var wg sync.WaitGroup
	for _, host := range task.Hosts {
		wg.Add(1)
		go func(h string) {
			defer wg.Done()
			color.PrintRunning("[%s][%s] Running...", task.Name, h)
			color.PrintRunning(renderedCmd)
			output, err := r.Exec.Execute(h, envAssign, renderedCmd)
			if err != nil {
				color.PrintError("[%s][%s] Error", task.Name, h)
				color.PrintError("%s", output)
			} else {
				color.PrintSuccess("[%s][%s] Success", task.Name, h)
				color.PrintSuccess("%s", output)
			}
		}(host)
	}

	wg.Wait()
	color.PrintHeader("=== Completed: %s ===\n", task.Name)
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