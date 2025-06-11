package main

import (
    "flag"
    "fmt"
    "io/ioutil"
    "log"
    "os/exec"
    "sync"

    "gopkg.in/yaml.v2"
)

// Task 定义单个任务的执行配置
type Task struct {
    Name  string   `yaml:"name"`
    Hosts []string `yaml:"hosts"`
    Cmd   string   `yaml:"cmd"`
}

// Config 包含多个 Task
type Config struct {
    Tasks []Task `yaml:"tasks"`
}

// 加载 YAML 配置
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

func main() {
    cfgPath := flag.String("config", "config.yaml", "配置文件路径")
    flag.Parse()

    cfg, err := loadConfig(*cfgPath)
    if err != nil {
        log.Fatalf("加载配置失败: %v", err)
    }

    // 逐任务执行
    for _, task := range cfg.Tasks {
        fmt.Printf("=== 开始任务: %s ===\n", task.Name)
        var wg sync.WaitGroup

        // 并发执行每个主机
        for _, host := range task.Hosts {
            wg.Add(1)
            go func(h string, cmdStr string) {
                defer wg.Done()
                cmd := exec.Command("ssh", h, cmdStr)
                output, err := cmd.CombinedOutput()
                if err != nil {
                    fmt.Printf("[%s][%s] 执行失败: %v\n输出: %s\n", task.Name, h, err, output)
                } else {
                    fmt.Printf("[%s][%s] 执行成功: %s\n", task.Name, h, output)
                }
            }(host, task.Cmd)
        }
        wg.Wait()
        fmt.Printf("=== 任务完成: %s ===\n\n", task.Name)
    }
}
