package main

import (
	"fmt"
	"mlang/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello, %s! Welcome to MLang REPL!\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
