package commands

import (
	"strconv"
	"strings"
)

type CommandHandler struct{}

func NewCommandsHandler() *CommandHandler {
	return &CommandHandler{}
}

func (c *CommandHandler) makeCommand(inp string) Command {
	inp = strings.TrimSpace(inp)
	args := strings.Split(inp, " ")

	switch args[0] {
	case "exit":
		exitCode, err := strconv.Atoi(args[1])
		if err != nil {
			exitCode = 1
		}
		return NewExitCommand(exitCode)
	case "echo":
		return NewEchoCommand(args[1:])
	default:
		return NewNotFoundCommand(args[0])
	}
}

func (c *CommandHandler) Handle(inp string) error {
	command := c.makeCommand(inp)
	return command.Run()
}

func (c *CommandHandler) cdCommand() error {
	return nil
}
