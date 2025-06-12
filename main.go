package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
	"sync"

	"gopkg.in/yaml.v2"
)

// 控制台输出颜色（护眼色）
const (
	AnsiReset   = "\033[0m"
	AnsiHeader  = "\033[0;36m"
	AnsiRunning = "\033[1;33m"
	AnsiSuccess = "\033[1;92m"
	AnsiError   = "\033[1;91m"
)

// RemoteTask 表示单个任务
type RemoteTask struct {
	Name  string            `yaml:"name"`
	Hosts []string          `yaml:"hosts"`
	Env   map[string]string `yaml:"env,omitempty"`
	Cmd   string            `yaml:"cmd"`
}

// TaskFile 表示整个 YAML 配置文件
type TaskFile struct {
	Env   map[string]string `yaml:"env,omitempty"`
	Tasks []RemoteTask      `yaml:"tasks"`
}

// loadTaskFile 读取并解析任务配置 YAML 文件
func loadTaskFile(path string) (*TaskFile, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var taskFile TaskFile
	if err := yaml.Unmarshal(data, &taskFile); err != nil {
		return nil, err
	}
	return &taskFile, nil
}

// buildEnvString 合并全局和任务级 env，构造形如 "KEY='val'" 的字符串
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

func main() {
	taskFilePath := flag.String("file", "task.yaml", "path to task file (default: task.yaml)")
	flag.Parse()
	selectedTaskNames := flag.Args()

	taskFile, err := loadTaskFile(*taskFilePath)
	if err != nil {
		log.Fatalf("failed to load task file: %v", err)
	}

	var tasksToRun []RemoteTask
	if len(selectedTaskNames) == 0 {
		tasksToRun = taskFile.Tasks
	} else {
		taskMap := make(map[string]RemoteTask)
		for _, t := range taskFile.Tasks {
			taskMap[t.Name] = t
		}
		for _, name := range selectedTaskNames {
			task, ok := taskMap[name]
			if !ok {
				log.Fatalf("task '%s' not found in task file", name)
			}
			tasksToRun = append(tasksToRun, task)
		}
	}

	for _, task := range tasksToRun {
		fmt.Printf("%s=== Task: %s ===%s\n", AnsiHeader, task.Name, AnsiReset)

		envAssign := buildEnvString(taskFile.Env, task.Env)

		commandBody := strings.TrimSpace(task.Cmd)
		if commandBody == "" {
			commandBody = "true"
		}

		var fullCommand string
		if envAssign != "" {
			fullCommand = fmt.Sprintf("env %s bash -c '%s'", envAssign, commandBody)
		} else {
			fullCommand = fmt.Sprintf("bash -c '%s'", commandBody)
		}

		var wg sync.WaitGroup
		for _, host := range task.Hosts {
			wg.Add(1)
			go func(h string) {
				defer wg.Done()
				fmt.Printf("%s[%s][%s] Running...%s\n", AnsiRunning, task.Name, h, AnsiReset)

				cmd := exec.Command("ssh", "-T", h, fullCommand)
				output, err := cmd.CombinedOutput()
				if err != nil {
					fmt.Printf("%s[%s][%s] Error%s\n", AnsiError, task.Name, h, AnsiReset)
					fmt.Printf("%s%s%s\n", AnsiError, string(output), AnsiReset)
				} else {
					fmt.Printf("%s[%s][%s] Success%s\n", AnsiSuccess, task.Name, h, AnsiReset)
					fmt.Printf("%s%s%s\n", AnsiSuccess, string(output), AnsiReset)
				}
			}(host)
		}
		wg.Wait()
		fmt.Printf("%s=== Completed: %s ===%s\n\n", AnsiHeader, task.Name, AnsiReset)
	}
}