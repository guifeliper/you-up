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
		fmt.Println("expected 'get' or 'auth' subcommands")
		os.Exit(1)
	}

	//look at the 2nd argument's value
	switch os.Args[1] {
	case "auth":
		controller.GoogleLogin()
	case "get":
		controller.GetChannelByUsername()
	case "upload":
		controller.UploadVideo()

	default: // if we don't understand the input
	}

}
