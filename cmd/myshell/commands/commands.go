package commands

import (
	"fmt"
	"os"
	"os/exec"
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

type ExternalCommand struct {
	args []string
}

func NewExternalCommand(args []string) *ExternalCommand {
	return &ExternalCommand{
		args: args,
	}
}

func (n *ExternalCommand) Run() error {
	command := exec.Command(n.args[0], n.args[1:]...)
	command.Stderr = os.Stderr
	command.Stdout = os.Stdout

	err := command.Run()
	if err != nil {
		fmt.Printf("%s: command not found\n", n.args[0])
	}
	return nil
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

	var buildIns = []string{"exit", "echo", "type", "pwd"}
	if contains(buildIns, command) {
		fmt.Printf("%s is a shell builtin\n", command)
		return nil
	}

	paths := strings.Split(os.Getenv("PATH"), ":")
	for _, path := range paths {
		if _, err := os.Stat(fmt.Sprintf("%s/%s", path, command)); err == nil {
			fmt.Printf("%s is %s/%s\n", command, path, command)
			return nil
		}
	}

	fmt.Printf("%s not found\n", command)
	return nil
}

type PwdCommand struct{}

func NewPwdCommand() *PwdCommand {
	return &PwdCommand{}
}

func (p *PwdCommand) Run() error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", wd)
	return nil
}

type CdCommand struct {
	args []string
}

func NewCdCommand(args []string) *CdCommand {
	return &CdCommand{
		args: args,
	}
}

func (c *CdCommand) Run() error {
	if len(c.args) == 1 {
		err := os.Chdir(c.args[0])
		if err != nil {
			return fmt.Errorf("%s: No such file or directory", c.args[0])
		}
	} else {
		return os.Chdir(os.Getenv("HOME"))
	}
	return nil
}
