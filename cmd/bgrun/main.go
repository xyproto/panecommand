package main

import (
	"fmt"
	"time"

	"github.com/xyproto/panecommand"
)

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func main() {
	pc, err := panecommand.New("bash", "-c", "for i in {1..20}; do echo $i; sleep 0.5; done")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	done := make(chan struct{})
	go func() {
		if err := pc.Run(); err != nil {
			fmt.Println("Command completed with error:", err)
		}
		close(done)
	}()

	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

loop:
	for {
		select {
		case <-done:
			break loop
		case <-ticker.C:
			clearScreen()
			for _, line := range pc.Lines {
				fmt.Println(line)
			}
		}
	}
}
