package main

import (
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
)

func askWithCheckboxes(message string, options []string) []string {
	userChoose := []string{}

	prompt := &survey.MultiSelect{
		Message:  message,
		Options:  options,
		PageSize: 5,
	}
	survey.AskOne(prompt, &userChoose)

	return userChoose
}

func main() {
	answers := askWithCheckboxes(
		"Which of following fruits you like?",
		[]string{
			"Apple",
			"Banana",
			"Orange",
			"Pineapple",
			"Grape",
			"Strawberry",
		},
	)

	favoriteFruits := strings.Join(answers, ", ")

	fmt.Println("Oh, I see! You like: ", favoriteFruits)
}
