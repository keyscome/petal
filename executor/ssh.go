// executor/ssh.go
package executor

import (
	"fmt"
	"os/exec"
	"strings"
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
