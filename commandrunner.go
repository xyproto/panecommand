package panecommand

import (
	"bytes"
	"os/exec"
	"time"
)

type CommandRunner struct {
	command string
	output  chan string
}

func NewCommandRunner(command string) *CommandRunner {
	return &CommandRunner{
		command: command,
		output:  make(chan string),
	}
}

func (cr *CommandRunner) Run() {
	for {
		cmd := exec.Command("sh", "-c", cr.command)
		var out bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &out
		err := cmd.Run()
		if err != nil {
			cr.output <- err.Error()
		} else {
			cr.output <- out.String()
		}
		time.Sleep(10 * time.Second) // sleep for N seconds
	}
}

func (cr *CommandRunner) Output() chan string {
	return cr.output
}
