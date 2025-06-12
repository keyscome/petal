// main.go
package main

import (
	"flag"
	"log"

	"keyscome.io/petal/config"
	"keyscome.io/petal/executor"
	"keyscome.io/petal/task"
)

func main() {
	taskFilePath := flag.String("file", "task.yaml", "path to task file (default: task.yaml)")
	flag.Parse()
	selectedTaskNames := flag.Args()

	taskFile, err := config.LoadTaskFile(*taskFilePath)
	if err != nil {
		log.Fatalf("failed to load task file: %v", err)
	}

	if err := config.ValidateTaskFile(taskFile); err != nil {
		log.Fatalf("invalid task file: %v", err)
	}

	runner := task.NewRunner(taskFile.Env, executor.SSHExecutor{})

	var tasksToRun []config.RemoteTask
	if len(selectedTaskNames) == 0 {
		tasksToRun = taskFile.Tasks
	} else {
		m := make(map[string]config.RemoteTask)
		for _, t := range taskFile.Tasks {
			m[t.Name] = t
		}
		for _, name := range selectedTaskNames {
			t, ok := m[name]
			if !ok {
				log.Fatalf("task '%s' not found in task file", name)
			}
			tasksToRun = append(tasksToRun, t)
		}
	}

	for _, t := range tasksToRun {
		runner.Run(t)
	}
}
