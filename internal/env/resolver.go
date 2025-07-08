// File: internal/env/resolver.go
package env

import (
	"bufio"
	"os"
	"regexp"
	"strings"
	"text/template"
)

// EnvVar represents a single environment variable with a value and a type (plain or secret)
type EnvVar struct {
	Value string
	Type  string // "plain" or "secret"
}

// Resolver loads and renders environment variables from global, file, and task scopes
type Resolver struct {
	merged map[string]EnvVar // task > file > global
}

// NewResolver constructs a Resolver given multiple scopes of environment variable maps
func NewResolver(global, file, task map[string]EnvVar) *Resolver {
	merged := make(map[string]EnvVar)
	for k, v := range global {
		merged[k] = v
	}
	for k, v := range file {
		merged[k] = v
	}
	for k, v := range task {
		merged[k] = v
	}
	r := &Resolver{merged: merged}
	r.preprocess()
	return r
}

func (r *Resolver) Get(key string) (string, bool) {
	if v, ok := r.merged[key]; ok {
		return v.Value, true
	}
	return "", false
}

func (r *Resolver) Flat() map[string]string {
	out := make(map[string]string)
	for k, v := range r.merged {
		out[k] = v.Value
	}
	return out
}

func (r *Resolver) Render(input string) string {
	// Handle ${VAR} syntax
	out := r.replaceVars(input)
	// Handle {{ .VAR }} syntax using text/template
	tmpl, err := template.New("env").Parse(out)
	if err != nil {
		return out
	}
	var sb strings.Builder
	tmpl.Execute(&sb, r.Flat())
	return sb.String()
}

func (r *Resolver) replaceVars(s string) string {
	varRegex := regexp.MustCompile(`\$\{(\w+)}`)
	return varRegex.ReplaceAllStringFunc(s, func(match string) string {
		key := varRegex.FindStringSubmatch(match)[1]
		if val, ok := r.Get(key); ok {
			return val
		}
		return match
	})
}

// preprocess does recursive variable replacement on values (up to 3 passes)
func (r *Resolver) preprocess() {
	for i := 0; i < 3; i++ {
		for k, v := range r.merged {
			resolved := r.replaceVars(v.Value)
			// avoid infinite loop
			if resolved != v.Value {
				r.merged[k] = EnvVar{Value: resolved, Type: v.Type}
			}
		}
	}
}

func (r *Resolver) GetMaskedEnv() map[string]string {
	masked := make(map[string]string)
	for k, v := range r.merged {
		if v.Type == "secret" {
			masked[k] = "******"
		} else {
			masked[k] = v.Value
		}
	}
	return masked
}

func (r *Resolver) CheckMissing(required []string) []string {
	missing := []string{}
	for _, k := range required {
		if _, ok := r.Get(k); !ok {
			missing = append(missing, k)
		}
	}
	return missing
}

func LoadEnvFile(path string) (map[string]EnvVar, error) {
	result := make(map[string]EnvVar)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		k := strings.TrimSpace(parts[0])
		v := strings.TrimSpace(parts[1])
		// result[k] = EnvVar{Value: v, Type: "plain"}
		result[k] = EnvVar{Value: v} // 默认 Type 留空
	}
	return result, scanner.Err()
}

func MergeFromStringMap(in map[string]string, asType string) map[string]EnvVar {
	res := make(map[string]EnvVar)
	for k, v := range in {
		res[k] = EnvVar{Value: v, Type: asType}
	}
	return res
}
