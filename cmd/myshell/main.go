package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Fprint(os.Stdout, "$ ")

	userInp, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return
	}

	handler := newCommandsHandler()

	err = handler.handle(userInp)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
