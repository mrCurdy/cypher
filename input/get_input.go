package input

import (
	"fmt"
)

func GetInput() (toEncrypt bool, encoding string, message string) {

	falseInputs := true

	welcomeSign := "Welcome to the Cypher Tool!"
	encDecMenu := "\nSelect operation (1/2):\n1. Encrypt.\n2. Decrypt."
	wrongChoiceSign := "\nWrong selection. Please try again."
	cypherSelectMenu := "\nSelect cypher (1/3):\n1. ROT13.\n2. Reverse.\n3. Reverse message and shift"
	entMessageSign := "Enter the message:"

	fmt.Println(welcomeSign)

	for falseInputs {
		var i int
		fmt.Println(encDecMenu)
		fmt.Scanln(&i)

		switch i {
		case 1:
			toEncrypt = true
		case 2:
			toEncrypt = false
		default:
			fmt.Println(wrongChoiceSign)
			continue
		}

		for message == "" {
			fmt.Println(cypherSelectMenu)
			encoding = Scanner()

			if encoding == "1" || encoding == "2" || encoding == "3" {

				fmt.Println(entMessageSign)
				message = Scanner()
				fmt.Println()
				falseInputs = false

			} else {
				fmt.Println(wrongChoiceSign)
				continue
			}
		}
	}
	return toEncrypt, encoding, message
}
