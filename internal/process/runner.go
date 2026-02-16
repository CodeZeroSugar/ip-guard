// Package process provides an os.exec wrapper as an interface for other packages to execute commands
package process

import "os/exec"

type Runner interface {
	Run(name string, args ...string) ([]byte, error)
}

type OSRunner struct{}

func (r OSRunner) Run(name string, args ...string) ([]byte, error) {
	cmd := exec.Command(name, args...)
	return cmd.CombinedOutput()
}
