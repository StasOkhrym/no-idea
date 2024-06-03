package main

import (
	"bufio"
	"fmt"
	"os"
)

type REPL struct {
	handler *commandsHandler
}

func NewREPL() *REPL {
	return &REPL{
		handler: newCommandsHandler(),
	}
}

func (r *REPL) Run() {
	for {
		fmt.Fprint(os.Stdout, "$ ")

		userInp, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			return
		}

		err = r.handler.handle(userInp)
		if err != nil {
			fmt.Println(err)
		}
	}
}
