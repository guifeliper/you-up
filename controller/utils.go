package controller

import (
	"log"
	"os/user"
)

func GetHomeDir() string {
	usr, err := user.Current()
	HandleError(err, "Home directory not available")
	return usr.HomeDir
}

func HandleError(err error, message string) {
	if message == "" {
		message = "Error making API call"
	}
	if err != nil {
		log.Fatalf(message+": %v", err.Error())
	}
}
