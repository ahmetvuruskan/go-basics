package main

import (
	"fmt"
	"os"
)

var users = [2][2]string{
	{"inanc", "1879"},
	{"jack", "1888"},
}

func main() {

	l := len(os.Args) - 1
	if l == 0 {
		fmt.Printf("Usage: [username] [password]")
		os.Exit(0)
	}
	validateUser(os.Args[1], os.Args[2])
}

func validateUser(username, password string) {

	const (
		invalidPassword = "Invalid password for %q \n"
		accessDenied    = "Access denied for %q \n"
		accessGranted   = "Access granted for %q \n"
	)

	var isUserExists bool = false
	var findUser [2]string

	for _, user := range users {
		if user[0] == username {
			isUserExists = true
			findUser = user
		}
	}

	if isUserExists {
		if findUser[1] == password {
			fmt.Printf(accessGranted, username)
		} else {
			fmt.Printf(invalidPassword, username)
		}
	} else {
		fmt.Printf(accessDenied, username)
		os.Exit(0)
	}

}
