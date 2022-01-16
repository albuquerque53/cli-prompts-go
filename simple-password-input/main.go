package main

import (
	"crypto/sha256"
	"fmt"
	"os"
	"syscall"

	"golang.org/x/term"
)

func AskPassword() string {
	var password string

	for {
		fmt.Fprint(os.Stderr, "Type your password: ")
		input, _ := term.ReadPassword(int(syscall.Stdin))

		password = string(input)

		if password != "" {
			break
		}

	}

	hashedPassword := sha256.Sum256([]byte(password))
	return string(hashedPassword[:])
}

func main() {
	password := AskPassword()

	fmt.Printf("\nOh, that's your password? %x\n", password)
}
