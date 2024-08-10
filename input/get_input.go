package input

import (
	"fmt"
)

func GetInput() (toEncrypt bool, encoding string, message string) {

	falseInputs := true

	fmt.Print("Welcome to the Cypher Tool!\n\nSelect operation (1/2):\n")

	for falseInputs {
		var i int
		fmt.Println("1. Encrypt.\n2. Decrypt.")
		fmt.Scanln(&i)

		switch i {
		case 1:
			toEncrypt = true
		case 2:
			toEncrypt = false
		default:
			fmt.Print("\nSelect correct operation 1 or 2\n")
			continue
		}

		fmt.Println("\nSelect cypher (1/3):\n1. ROT13.\n2. Reverse.\n3. Reverse message and shift")
		fmt.Scanln(&encoding)

		if encoding == "1" || encoding == "2" || encoding == "3" {

			fmt.Println("\nEnter the message:")
			message = Scanner()
			fmt.Println()
			falseInputs = false

		} else {
			continue
		}
	}
	return toEncrypt, encoding, message
}
