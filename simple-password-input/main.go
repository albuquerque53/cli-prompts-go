package main

import (
	"fmt"
	"os"
	"syscall"

	"golang.org/x/term"
)

func AskPassword() string {
	var password string

	for {
		fmt.Fprint(os.Stderr, "Type your password (will be printed on output): ")
		input, _ := term.ReadPassword(int(syscall.Stdin))

		password = string(input)

		if password != "" {
			break
		}
	}

	return password
}

func main() {
	password := AskPassword()

	fmt.Printf("\nOh, that's your password? %s\n", password)
}
