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

// 颜色常量（护眼配色）
const (
    ColorReset   = "\033[0m"
    ColorHeader  = "\033[0;36m" // Cyan: task headers
    ColorRun     = "\033[1;33m" // Bold Yellow: running
    ColorSuccess = "\033[1;92m" // Bright Green: success
    ColorError   = "\033[1;91m" // Bright Red: error
)

// Task 定义
type Task struct {
    Name  string            `yaml:"name"`
    Hosts []string          `yaml:"hosts"`
    Env   map[string]string `yaml:"env,omitempty"`
    Cmd   string            `yaml:"cmd"`
}

// Config 定义
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

// 拼出 "KEY='val'" 串，不带尾空格
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
    // 按空格拼接： "VAR1='v1' VAR2='v2'"
    return strings.Join(parts, " ")
}

func main() {
    cfgFile := flag.String("config", "config.yaml", "path to config")
    flag.Parse()

    cfg, err := loadConfig(*cfgFile)
    if err != nil {
        log.Fatalf("load config: %v", err)
    }

    for _, task := range cfg.Tasks {
        fmt.Printf("%s=== Task: %s ===%s\n", ColorHeader, task.Name, ColorReset)

        // 1) 构造 assignment 前缀
        assign := buildAssignPrefix(cfg.Env, task.Env) // e.g. "FAVORITE_CAR='Mazda'"

        // 2) 构造 echo 语句
        echoStmt := ""
        if assign != "" {
            for _, kv := range strings.Split(assign, " ") {
                key := strings.SplitN(kv, "=", 2)[0]
                echoStmt += fmt.Sprintf("echo %s=$%s; ", key, key)
            }
        }

        // 3) 业务命令或默认 true
        body := strings.TrimSpace(task.Cmd)
        if body == "" {
            body = "true"
        }
        // 支持多行块：保持原有 task.Cmd 里的换行
        // wrappedCmd = "VAR=val VAR2=val2 echo ...; your_cmd"
        var wrappedCmd string
        if assign != "" {
            wrappedCmd = assign + " " + echoStmt + body
        } else {
            wrappedCmd = echoStmt + body
        }

        var wg sync.WaitGroup
        for _, host := range task.Hosts {
            wg.Add(1)
            go func(h string) {
                defer wg.Done()
                fmt.Printf("%s[%s][%s] Running...%s\n",
                    ColorRun, task.Name, h, ColorReset)

                // bash -lc 保留 profile，然后执行 assignment+echo+body
                cmd := exec.Command("ssh", "-T", h,
                    "bash", "-lc", wrappedCmd)
                out, err := cmd.CombinedOutput()

                if err != nil {
                    fmt.Printf("%s[%s][%s] Error%s\n",
                        ColorError, task.Name, h, ColorReset)
                    fmt.Printf("%s%s%s\n", ColorError, string(out), ColorReset)
                } else {
                    fmt.Printf("%s[%s][%s] Success%s\n",
                        ColorSuccess, task.Name, h, ColorReset)
                    fmt.Printf("%s%s%s\n", ColorSuccess, string(out), ColorReset)
                }
            }(host)
        }
        wg.Wait()
        fmt.Printf("%s=== Completed: %s ===%s\n\n",
            ColorHeader, task.Name, ColorReset)
    }
}