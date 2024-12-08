package main

import (
	"bytes"
	"log"

	"github.com/gdamore/tcell"
	"github.com/jkimunyi-dev/terminal-screen-renderer/internal/renderer"
	"github.com/jkimunyi-dev/terminal-screen-renderer/internal/stream"
)

func main() {
	// Create renderer
	term, err := renderer.NewTerminalRenderer()
	if err != nil {
		log.Fatalf("Failed to create renderer: %v", err)
	}
	defer term.Close()

	// Example binary stream
	exampleStream := []byte{
		// Screen setup: 80x24, 16 colors
		0x1, 80, 24, 0x01,

		// Draw character: 'H' at (10,5) in red
		0x2, 10, 5, 1, 'H',

		// Draw another character: 'i' at (11,5) in green
		0x2, 11, 5, 2, 'i',

		// End of stream
		0xFF,
	}

	// Create a byte reader
	reader := bytes.NewReader(exampleStream)
	parser := stream.NewParser(reader)

	// Process commands
	for {
		cmd, err := parser.ParseNextCommand()
		if err != nil {
			log.Fatalf("Error parsing command: %v", err)
		}

		// Nil command indicates end of stream
		if cmd == nil {
			break
		}

		if err := term.ProcessCommand(cmd); err != nil {
			log.Fatalf("Error processing command: %v", err)
		}
	}

	// Keep screen open until 'q' is pressed
	for {
		ev := term.screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			if ev.Rune() == 'q' {
				return
			}
		}
	}
}
