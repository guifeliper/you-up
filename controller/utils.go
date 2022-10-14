package controller

import (
	"os/user"
)

func GetHomeDir() string {
	usr, err := user.Current()
	handleError(err, "Home directory not available")
	return usr.HomeDir
}
