// task/runner.go
package task

import (
	"fmt"
	"strings"
	"sync"

	"keyscome.io/petal/color"
	"keyscome.io/petal/config"
)

type Executor interface {
	Execute(host, envAssign, cmd string) (string, error)
}

type Runner struct {
	GlobalEnv map[string]string
	Exec      Executor
}

func NewRunner(globalEnv map[string]string, exec Executor) *Runner {
	return &Runner{
		GlobalEnv: globalEnv,
		Exec:      exec,
	}
}

func (r *Runner) Run(task config.RemoteTask) {
	color.PrintHeader("=== Task: %s ===", task.Name)

	envAssign := buildEnvString(r.GlobalEnv, task.Env)
	var wg sync.WaitGroup

	for _, host := range task.Hosts {
		wg.Add(1)
		go func(h string) {
			defer wg.Done()
			color.PrintRunning("[%s][%s] Running...", task.Name, h)
			color.PrintRunning(task.Cmd)
			output, err := r.Exec.Execute(h, envAssign, task.Cmd)
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

func buildEnvString(globalEnv, taskEnv map[string]string) string {
	merged := make(map[string]string)
	for k, v := range globalEnv {
		merged[k] = v
	}
	for k, v := range taskEnv {
		merged[k] = v
	}
	if len(merged) == 0 {
		return ""
	}
	parts := make([]string, 0, len(merged))
	for k, v := range merged {
		parts = append(parts, fmt.Sprintf("%s='%s'", k, v))
	}
	return strings.Join(parts, " ")
}
