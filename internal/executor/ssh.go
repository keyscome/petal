// File: internal/executor/executor.go
package executor

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"keyscome.io/petal/internal/env"
	"keyscome.io/petal/internal/config"
)

type SSHExecutor struct{}

func (s SSHExecutor) Execute(host, envAssign, cmd string) (string, error) {
	cmdBody := strings.TrimSpace(cmd)
	if cmdBody == "" {
		cmdBody = "true"
	}

	// 构建远程 bash 执行命令
	var sshArgs []string
	sshArgs = append(sshArgs, "-T", host)

	// 加上 env 变量（如果有）和 bash -s
	var remoteCommand string
	if envAssign != "" {
		remoteCommand = fmt.Sprintf("env %s bash -s", envAssign)
	} else {
		remoteCommand = "bash -s"
	}
	sshArgs = append(sshArgs, remoteCommand)

	// 构造命令
	sshCmd := exec.Command("ssh", sshArgs...)

	// 将命令内容通过 stdin 传入
	sshCmd.Stdin = strings.NewReader(cmdBody)

	// 获取输出
	output, err := sshCmd.CombinedOutput()
	return string(output), err
}

// Run executes a RemoteTask with environment resolution
func Run(t config.RemoteTask, globalEnv map[string]env.EnvVar, envFilePath string) error {
	// Load env from file
	fileEnv := map[string]env.EnvVar{}
	if envFilePath != "" {
		loaded, err := env.LoadEnvFile(envFilePath)
		if err != nil {
			return fmt.Errorf("failed to load env file: %w", err)
		}
		fileEnv = loaded
	}

	// Convert task.Env from map[string]string to map[string]EnvVar
	taskEnv := env.MergeFromStringMap(t.Env)

	// Build resolver
	resolver := env.NewResolver(globalEnv, fileEnv, taskEnv)

	// Check for required vars (optional)
	// missing := resolver.CheckMissing([]string{"TMP_DIR", "NODE1_IP"})
	// if len(missing) > 0 {
	// 	return fmt.Errorf("missing env vars: %v", missing)
	// }

	// Render command
	renderedCmd := resolver.Render(t.Cmd)
	fmt.Printf("[Task: %s] Command: %s\n", t.Name, renderedCmd)

	// Prepare command
	cmd := exec.Command("bash", "-", renderedCmd)

	// Inject merged env
	envMap := resolver.Flat()
	envList := []string{}
	for k, v := range envMap {
			envList = append(envList, fmt.Sprintf("%s=%s", k, v))
	}

	cmd.Env = append(os.Environ(), envList...) 

	// Run and stream output
	output, err := cmd.CombinedOutput()
	fmt.Printf("[Task: %s] Output:\n%s\n", t.Name, string(output))
	if err != nil {
		return fmt.Errorf("command failed: %w", err)
	}

	return nil
}