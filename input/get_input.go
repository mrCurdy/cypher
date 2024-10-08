package input

import (
	"fmt"
)

func GetInput() (toEncrypt bool, encoding string, message string) {

	enteringMessage := true
	enteringFirstInput := true
	enteringSecondInput := true
	firstInput := ""
	secondInput := ""

	welcomeSign := "Welcome to the Cypher Tool!"
	encDecMenu := "\nSelect operation (1/2):\n1. Encrypt.\n2. Decrypt."
	wrongChoiceSign := "\nWrong selection. Please try again."
	cypherSelectMenu := "\nSelect cypher (1/3):\n1. ROT13.\n2. Reverse.\n3. Reverse and shift"
	entMessageSign := "\nEnter the message:"
	wrongMessageSign := "\nMessage is empty. Try again"

	fmt.Println(welcomeSign)

	for enteringFirstInput {

		fmt.Println(encDecMenu)
		firstInput = Scanner()

		switch firstInput {
		case "1":
			toEncrypt = true
		case "2":
			toEncrypt = false
		default:
			fmt.Println(wrongChoiceSign)
			continue
		}
		enteringFirstInput = false
	}

	for enteringSecondInput {

		fmt.Println(cypherSelectMenu)
		secondInput = Scanner()

		switch secondInput {
		case "1":
			encoding = "1"
		case "2":
			encoding = "2"
		case "3":
			encoding = "3"
		default:
			fmt.Println(wrongChoiceSign)
			continue
		}
		enteringSecondInput = false
	}

	for enteringMessage {

		fmt.Println(entMessageSign)
		message = Scanner()
		if len(message) == 0 || message == " " {
			fmt.Println(wrongMessageSign)
			continue
		}
		enteringMessage = false

	}

	return toEncrypt, encoding, message
}
