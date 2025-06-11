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

// ANSI color codes (subtle palette)
const (
    ColorReset = "\033[0m"
    ColorGray  = "\033[0;37m"
    ColorCyan  = "\033[0;36m"
    ColorGreen = "\033[0;32m"
    ColorRed   = "\033[0;31m"
)

// Task defines a single remote job
type Task struct {
    Name  string   `yaml:"name"`
    Hosts []string `yaml:"hosts"`
    Cmd   string   `yaml:"cmd"`
}

// Config holds all tasks
type Config struct {
    Tasks []Task   `yaml:"tasks"`
}

// reads YAML configuration
func loadConfig (path string) (*Config, error) {
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
    // Load configuration
    cfgPath := flag.String("config", "config.yaml", "path to config file")
    flag.Parse()

    cfg, err := loadConfig(*cfgPath)
    if err != nil {
        log.Fatalf("Failed to load config: %v", err)
    }

    // Iterate tasks
    for _, task := range cfg.Tasks {
        // Task header in cyan
        fmt.Printf("%s=== Task: %s ===%s\n", ColorCyan, task.Name, ColorReset)
        var wg sync.WaitGroup

        for _, host := range task.Hosts {
            wg.Add(1)
            go func(h string) {
                defer wg.Done()
                // Label in gray
                fmt.Printf("%s[%s][%s] Running...%s\n", ColorGray, task.Name, h, ColorReset)

                cmd := exec.Command("ssh", h, task.Cmd)
                output, err := cmd.CombinedOutput()

                if err != nil {
                    // Error in red
                    fmt.Printf("%s[%s][%s] Error%s\n", ColorRed, task.Name, h, ColorReset)
                    // Print detailed output
                    fmt.Printf("%s%s%s\n", ColorRed, string(output), ColorReset)
                } else {
                    // Success in green
                    fmt.Printf("%s[%s][%s] Success%s\n", ColorGreen, task.Name, h, ColorReset)
                    fmt.Printf("%s%s%s\n", ColorGreen, string(output), ColorReset)
                }
            }(host)
        }

        wg.Wait()
        // Completion in cyan
        fmt.Printf("%s=== Completed: %s ===%s\n\n", ColorCyan, task.Name, ColorReset)
    }
}