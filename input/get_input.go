package input

import (
	"fmt"
)

func GetInput() (toEncrypt bool, encoding string, message string) {

	falseInputs := true

	welcomeSign := "Welcome to the Cypher Tool!\n\nSelect operation (1/2):\n"
	encDecMenu := "1. Encrypt.\n2. Decrypt."
	wrongChoiceSign := "\nSelect correct operation 1 or 2\n"
	cypherSelectMenu := "\nSelect cypher (1/3):\n1. ROT13.\n2. Reverse.\n3. Reverse message and shift"
	entMessageSign := "\nEnter the message:"

	fmt.Print(welcomeSign)

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
			fmt.Print(wrongChoiceSign)
			continue
		}

		fmt.Println(cypherSelectMenu)
		encoding = Scanner()

		if encoding == "1" || encoding == "2" || encoding == "3" {

			fmt.Println(entMessageSign)
			message = Scanner()
			fmt.Println()
			falseInputs = false

		} else {
			continue
		}
	}
	return toEncrypt, encoding, message
}
