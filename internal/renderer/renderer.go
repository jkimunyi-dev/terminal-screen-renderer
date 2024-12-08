package renderer

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/jkimunyi-dev/terminal-screen-renderer/internal/stream"
)

type TerminalRenderer struct {
	screen    tcell.Screen
	width     int
	height    int
	colorMode ColorMode
	cursorX   int
	cursorY   int
}

// NewTerminalRenderer initializes a new terminal renderer
func NewTerminalRenderer() (*TerminalRenderer, error) {
	// Initialize tcell screen
	s, err := tcell.NewScreen()
	if err != nil {
		return nil, fmt.Errorf("failed to create screen: %v", err)
	}
	if err := s.Init(); err != nil {
		return nil, fmt.Errorf("failed to initialize screen: %v", err)
	}

	return &TerminalRenderer{
		screen: s,
	}, nil
}

// ProcessCommand handles individual rendering commands
func (r *TerminalRenderer) ProcessCommand(cmd *stream.Command) error {
	switch cmd.Type {
	case stream.CommandScreenSetup:
		return r.setupScreen(cmd.Data)
	case stream.CommandDrawCharacter:
		return r.drawCharacter(cmd.Data)
	// Other command handlers would be similar
	default:
		return fmt.Errorf("unhandled command type: %x", cmd.Type)
	}
}

// setupScreen initializes the terminal screen
func (r *TerminalRenderer) setupScreen(data []byte) error {
	if len(data) < 3 {
		return fmt.Errorf("insufficient data for screen setup")
	}

	r.width = int(data[0])
	r.height = int(data[1])
	r.colorMode = ColorMode(data[2])

	// Set screen size
	r.screen.SetSize(r.width, r.height)
	r.screen.Clear()

	return nil
}

// drawCharacter renders a single character
func (r *TerminalRenderer) drawCharacter(data []byte) error {
	if len(data) < 4 {
		return fmt.Errorf("insufficient data for drawing character")
	}

	x := int(data[0])
	y := int(data[1])
	colorIndex := data[2]
	char := rune(data[3])

	color := MapColor(r.colorMode, colorIndex)
	style := tcell.StyleDefault.Foreground(color)

	r.screen.SetContent(x, y, char, nil, style)
	r.screen.Show()

	return nil
}

// Close finalizes the screen
func (r *TerminalRenderer) Close() {
	r.screen.Fini()
}
