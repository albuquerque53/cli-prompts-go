package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func AskName() string {
	var name string
	input := bufio.NewReader(os.Stdin)

	for {
		fmt.Fprint(os.Stderr, "Whats's your name? ")
		name, _ = input.ReadString('\n')

		if name != "" {
			break
		}
	}

	return strings.TrimSpace(name)
}

func main() {
	name := AskName()

	fmt.Printf("Hello, %s\n", name)
}
