package main

import (
	"flag"
	"fmt"
	"os"

	"you-up/controller"
)

func main() {

	//'you-up auth' subcommand
	flag.NewFlagSet("auth", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Println("expected 'get' or 'add' subcommands")
		os.Exit(1)
	}

	//look at the 2nd argument's value
	switch os.Args[1] {
	case "auth": // if its the 'get' command
		controller.GoogleLogin()
	default: // if we don't understand the input
	}

}
