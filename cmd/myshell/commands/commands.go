package commands

import (
	"fmt"
	"os"
	"strings"
)

type Command interface {
	Run() error
}

type ExitCommand struct {
	exitCode int
}

func NewExitCommand(exitCode int) *ExitCommand {
	return &ExitCommand{
		exitCode: exitCode,
	}
}

func (e *ExitCommand) Run() error {
	os.Exit(e.exitCode)
	return nil
}

type NotFoundCommand struct {
	command string
}

func NewNotFoundCommand(command string) *NotFoundCommand {
	return &NotFoundCommand{
		command: command,
	}
}

func (n *NotFoundCommand) Run() error {
	return fmt.Errorf("%s: command not found", n.command)
}

type EchoCommand struct {
	args []string
}

func NewEchoCommand(args []string) *EchoCommand {
	return &EchoCommand{
		args: args,
	}
}

func (e *EchoCommand) Run() error {
	echoString := fmt.Sprintf("%s", strings.Join(e.args, " "))
	fmt.Println(echoString)
	return nil
}
