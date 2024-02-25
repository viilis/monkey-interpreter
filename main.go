package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/viilis/repl"
)

func main() {
	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is Monkey programming language!\n", user.Username)
	fmt.Printf("Feel free to type in command\n")
	repl.StartRepl(os.Stdin, os.Stdout)
}
