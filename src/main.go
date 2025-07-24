package main

import (
	"fmt"
	"os"
	"os/user"

	"itstejas.com/monkey-go/src/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! Welcome to the Monkey Language!\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")

	repl.Start(os.Stdin, os.Stdout)
}
