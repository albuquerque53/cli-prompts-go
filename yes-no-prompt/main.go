package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func AskQuestion(question string, defaultChoice bool) bool {
	possibleChoices := "Y/n"

	if !defaultChoice {
		possibleChoices = "y/N"
	}

	input := bufio.NewReader(os.Stdin)
	var userChoice string

	for {
		fmt.Fprintf(os.Stderr, "%s (%s) ", question, possibleChoices)

		userChoice, _ = input.ReadString('\n')
		userChoice = strings.TrimSpace(userChoice)

		if userChoice == "" {
			return defaultChoice
		}

		lowerCaseChoice := strings.ToLower(userChoice)

		if lowerCaseChoice == "y" || lowerCaseChoice == "yes" {
			return true
		}

		if lowerCaseChoice == "n" || lowerCaseChoice == "no" {
			return false
		}
	}
}

func main() {
	isGoNice := AskQuestion("Is Go a nice language?", true)

	if isGoNice {
		fmt.Println("Agree!")
	} else {
		fmt.Println("Huh?")
	}
}
