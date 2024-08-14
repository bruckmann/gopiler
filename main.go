package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/bruckmann/gopiler/console"
)

func main() {

	usr, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s, Welcome to the Summer Language console\n", usr.Username)
	// !TODO: Create an help command to help with the syntax
	fmt.Printf("If you need some help, try to use the command -h\n")

	console.Start(os.Stdin, os.Stdout)

}
