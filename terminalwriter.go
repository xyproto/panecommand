package panecommand

import (
	"fmt"
)

type TerminalWriter struct{}

func NewTerminalWriter() *TerminalWriter {
	return &TerminalWriter{}
}

func (tw *TerminalWriter) Write(x, y int, text string) {
	fmt.Printf("\033[%d;%dH%s", y, x, text)
}
