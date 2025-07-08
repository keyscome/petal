// File: main.go
package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"keyscome.io/petal/internal/config"
	"keyscome.io/petal/internal/env"
	"keyscome.io/petal/internal/executor"
	"keyscome.io/petal/internal/task"
)

func main() {
	// Define flags
	taskFilePath := flag.String("f", "task.yml", "task file path (e.g. -f task.yml)")
	envFilePath := flag.String("e", ".env", "env file path (e.g. -e .env)")
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [options] [task names...]\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
	selectedTaskNames := flag.Args()

	// Load task file
	taskFile, err := config.LoadTaskFile(*taskFilePath)
	if err != nil {
		log.Fatalf("failed to load task file: %v", err)
	}
	if err := config.ValidateTaskFile(taskFile); err != nil {
		log.Fatalf("invalid task file: %v", err)
	}

	// Merge global env
	globalEnv := env.MergeFromStringMap(taskFile.Env, "plain")

	// Init runner with global env and env file path
	runner := task.NewRunner(globalEnv, *envFilePath, executor.SSHExecutor{})

	// Select tasks
	var tasksToRun []config.RemoteTask
	if len(selectedTaskNames) == 0 {
		tasksToRun = taskFile.Tasks
	} else {
		taskMap := make(map[string]config.RemoteTask)
		for _, t := range taskFile.Tasks {
			taskMap[t.Name] = t
		}
		for _, name := range selectedTaskNames {
			t, ok := taskMap[name]
			if !ok {
				log.Fatalf("task '%s' not found in task file", name)
			}
			tasksToRun = append(tasksToRun, t)
		}
	}

	// Run selected tasks
	for _, t := range tasksToRun {
		runner.Run(t)
	}
}
