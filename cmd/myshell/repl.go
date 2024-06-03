package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/commands"
)

type REPL struct {
	handler *commands.CommandHandler
}

func NewREPL() *REPL {
	return &REPL{
		handler: commands.NewCommandsHandler(),
	}
}

func (r *REPL) Run() {
	for {
		fmt.Fprint(os.Stdout, "$ ")

		userInp, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			return
		}

		err = r.handler.Handle(userInp)
		if err != nil {
			fmt.Println(err)
		}
	}
}
