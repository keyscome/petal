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

	var fullCommand string
	if envAssign != "" {
		fullCommand = fmt.Sprintf("env %s bash -c '%s'", envAssign, cmdBody)
	} else {
		fullCommand = fmt.Sprintf("bash -c '%s'", cmdBody)
	}

	sshCmd := exec.Command("ssh", "-T", host, fullCommand)
	output, err := sshCmd.CombinedOutput()
	return string(output), err
}
