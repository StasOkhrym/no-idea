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

type TypeCommand struct {
	args []string
}

func NewTypeCommand(args []string) *TypeCommand {
	return &TypeCommand{
		args: args,
	}
}

func (t *TypeCommand) Run() error {
	command := t.args[0]

	var buildIns = []string{"exit", "echo", "type"}
	if contains(buildIns, command) {
		fmt.Printf("%s is a shell builtin\n", command)
	} else {
		fmt.Printf("%s not found\n", command)
	}

	return nil
}
