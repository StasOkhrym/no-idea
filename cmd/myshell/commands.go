package main

import (
	"fmt"
	"strings"
)

type commandsHandler struct{}

func newCommandsHandler() *commandsHandler {
	return &commandsHandler{}
}

func (c *commandsHandler) handle(inp string) error {
	str := strings.Trim(inp, "\n")
	switch str {
	case "cd":
		return c.cdCommand()
	default:
		return fmt.Errorf("%s: command not found", str)
	}
}

func (c *commandsHandler) cdCommand() error {
	return nil
}
