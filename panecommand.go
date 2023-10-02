package panecommand

import (
	"bufio"
	"io"
	"os/exec"
	"sync"
)

type PaneCommand struct {
	Cmd    *exec.Cmd
	Lines  []string
	Output io.ReadCloser
}

func New(command string, args ...string) (*PaneCommand, error) {
	cmd := exec.Command(command, args...)
	output, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	return &PaneCommand{Cmd: cmd, Output: output}, nil
}

func (pc *PaneCommand) Run() error {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		scanner := bufio.NewScanner(pc.Output)
		for scanner.Scan() {
			pc.Lines = append(pc.Lines, scanner.Text())
			if len(pc.Lines) > 5 {
				pc.Lines = pc.Lines[1:]
			}
		}
	}()

	if err := pc.Cmd.Start(); err != nil {
		return err
	}

	wg.Wait()
	return pc.Cmd.Wait()
}
