package gtools

import (
	"os/exec"
)

type cmdClient struct {
	OnWatch func(out string)
}

func (c *cmdClient) Write(p []byte) (int, error) {
	if c.OnWatch != nil {
		c.OnWatch(string(p))
	}
	return len(p), nil
}

func NewCmdClient(onWatch func(out string)) *cmdClient {
	return &cmdClient{
		OnWatch: onWatch,
	}
}

func (c *cmdClient) Command(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = c
	cmd.Stderr = c
	return cmd.Run()
}
