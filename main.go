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

const (
	ColorReset   = "\033[0m"
	ColorHeader  = "\033[0;36m"
	ColorRun     = "\033[1;33m"
	ColorSuccess = "\033[1;92m"
	ColorError   = "\033[1;91m"
)

type Task struct {
	Name  string            `yaml:"name"`
	Hosts []string          `yaml:"hosts"`
	Env   map[string]string `yaml:"env,omitempty"`
	Cmd   string            `yaml:"cmd"`
}

type Config struct {
	Env   map[string]string `yaml:"env,omitempty"`
	Tasks []Task            `yaml:"tasks"`
}

func loadConfig(path string) (*Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func buildAssignPrefix(globalEnv, taskEnv map[string]string) string {
	m := make(map[string]string)
	for k, v := range globalEnv {
		m[k] = v
	}
	for k, v := range taskEnv {
		m[k] = v
	}
	if len(m) == 0 {
		return ""
	}
	parts := make([]string, 0, len(m))
	for k, v := range m {
		parts = append(parts, fmt.Sprintf("%s='%s'", k, v))
	}
	return strings.Join(parts, " ")
}

func main() {
	cfgFile := flag.String("config", "task.yaml", "path to config file (default: task.yaml)")
	flag.Parse()
	taskNames := flag.Args() // 剩下的参数作为任务名顺序执行

	cfg, err := loadConfig(*cfgFile)
	if err != nil {
		log.Fatalf("load config: %v", err)
	}

	tasksToRun := make([]Task, 0)
	if len(taskNames) == 0 {
		tasksToRun = cfg.Tasks // 默认全部执行
	} else {
		// 按参数顺序筛选任务
		taskMap := make(map[string]Task)
		for _, t := range cfg.Tasks {
			taskMap[t.Name] = t
		}
		for _, name := range taskNames {
			task, ok := taskMap[name]
			if !ok {
				log.Fatalf("task '%s' not found in config", name)
			}
			tasksToRun = append(tasksToRun, task)
		}
	}

	for _, task := range tasksToRun {
		fmt.Printf("%s=== Task: %s ===%s\n", ColorHeader, task.Name, ColorReset)

		assign := buildAssignPrefix(cfg.Env, task.Env)
		body := strings.TrimSpace(task.Cmd)
		if body == "" {
			body = "true"
		}

		var wrappedCmd string
		if assign != "" {
			wrappedCmd = fmt.Sprintf("env %s bash -c '%s'", assign, body)
		} else {
			wrappedCmd = fmt.Sprintf("bash -c '%s'", body)
		}

		var wg sync.WaitGroup
		for _, host := range task.Hosts {
			wg.Add(1)
			go func(h string) {
				defer wg.Done()
				fmt.Printf("%s[%s][%s] Running...%s\n", ColorRun, task.Name, h, ColorReset)
				cmd := exec.Command("ssh", "-T", h, wrappedCmd)
				out, err := cmd.CombinedOutput()
				if err != nil {
					fmt.Printf("%s[%s][%s] Error%s\n", ColorError, task.Name, h, ColorReset)
					fmt.Printf("%s%s%s\n", ColorError, string(out), ColorReset)
				} else {
					fmt.Printf("%s[%s][%s] Success%s\n", ColorSuccess, task.Name, h, ColorReset)
					fmt.Printf("%s%s%s\n", ColorSuccess, string(out), ColorReset)
				}
			}(host)
		}
		wg.Wait()
		fmt.Printf("%s=== Completed: %s ===%s\n\n", ColorHeader, task.Name, ColorReset)
	}
}