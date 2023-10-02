package panecommand

import (
	"strings"
)

type OutputProcessor struct {
	lines  int
	length int
}

func NewOutputProcessor(lines, length int) *OutputProcessor {
	return &OutputProcessor{
		lines:  lines,
		length: length,
	}
}

func (op *OutputProcessor) Process(output string) string {
	lines := strings.Split(output, "\n")
	if len(lines) > op.lines {
		lines = lines[len(lines)-op.lines:]
	}
	for i, line := range lines {
		if len(line) > op.length {
			lines[i] = line[:op.length]
		}
	}
	return strings.Join(lines, "\n")
}
