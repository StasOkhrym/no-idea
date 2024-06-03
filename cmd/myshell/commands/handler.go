package commands

import (
	"strconv"
	"strings"
)

type CommandHandler struct{}

func NewCommandsHandler() *CommandHandler {
	return &CommandHandler{}
}

func (c *CommandHandler) makeCommand(args []string) Command {
	switch args[0] {
	case "exit":
		exitCode, err := strconv.Atoi(args[1])
		if err != nil {
			exitCode = 1
		}
		return NewExitCommand(exitCode)
	case "echo":
		return NewEchoCommand(args[1:])
	case "type":
		return NewTypeCommand(args[1:])
	case "pwd":
		return NewPwdCommand()
	case "cd":
		return NewCdCommand(args[1:])
	default:
		return NewExternalCommand(args)
	}
}

func (c *CommandHandler) Handle(inp string) error {
	inp = strings.TrimSpace(inp)
	args := strings.Split(inp, " ")

	command := c.makeCommand(args)
	return command.Run()
}
