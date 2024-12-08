package stream

import (
	"fmt"
	"io"
)

// CommandType represents different types of rendering commands
type CommandType byte

const (
	CommandScreenSetup   CommandType = 0x1
	CommandDrawCharacter CommandType = 0x2
	CommandDrawLine      CommandType = 0x3
	CommandRenderText    CommandType = 0x4
	CommandCursorMove    CommandType = 0x5
	CommandDrawAtCursor  CommandType = 0x6
	CommandClearScreen   CommandType = 0x7
	CommandEndOfStream   CommandType = 0xFF
)

// Command represents a parsed rendering command
type Command struct {
	Type CommandType
	Data []byte
}

// Parser handles parsing of the binary stream
type Parser struct {
	stream io.Reader
}

// NewParser creates a new stream parser
func NewParser(reader io.Reader) *Parser {
	return &Parser{stream: reader}
}

// ParseNextCommand reads and parses the next command from the stream
func (p *Parser) ParseNextCommand() (*Command, error) {
	// Read command byte
	cmdByte := make([]byte, 1)
	_, err := p.stream.Read(cmdByte)
	if err != nil {
		if err == io.EOF {
			return nil, nil
		}
		return nil, fmt.Errorf("error reading command byte: %v", err)
	}

	cmd := &Command{Type: CommandType(cmdByte[0])}

	// Determine data length based on command type
	switch cmd.Type {
	case CommandScreenSetup:
		cmd.Data = make([]byte, 3)
		_, err = io.ReadFull(p.stream, cmd.Data)
	case CommandDrawCharacter:
		cmd.Data = make([]byte, 4)
		_, err = io.ReadFull(p.stream, cmd.Data)
	case CommandDrawLine:
		cmd.Data = make([]byte, 6)
		_, err = io.ReadFull(p.stream, cmd.Data)
	case CommandCursorMove:
		cmd.Data = make([]byte, 2)
		_, err = io.ReadFull(p.stream, cmd.Data)
	case CommandDrawAtCursor:
		cmd.Data = make([]byte, 2)
		_, err = io.ReadFull(p.stream, cmd.Data)
	case CommandClearScreen, CommandEndOfStream:
		// No additional data
		cmd.Data = nil
	default:
		return nil, fmt.Errorf("unknown command type: 0x%x", cmd.Type)
	}

	if err != nil {
		return nil, fmt.Errorf("error reading command data: %v", err)
	}

	return cmd, nil
}
