package input

import (
	"fmt"
)

func getInput() (toEncrypt bool, encoding string, message string) {
	correctInputs := false
	for correctInputs {
		var i int
		fmt.Print("Welcome to the Cypher Tool!\n\nSelect operation (1/2):\n1. Encrypt.\n2. Decrypt.\n")
		fmt.Scanln(&i)
		fmt.Println()
		switch i {
		case 1:
			toEncrypt = true
		case 2:
			toEncrypt = false
		default:
			fmt.Print("Select correct operation 1 or 2")
			continue
		}
		fmt.Println("Select cypher (1/2):\n1. ROT13.\n2. Reverse.\n3. Reverse message and shift")
		fmt.Scanln(&encoding)
		fmt.Println()
		if encoding == "1" || encoding == "2" || encoding == "3" {
			fmt.Println("Enter the message:")
			fmt.Scanln(&message)
		} else {
			continue
		}
	}
	return toEncrypt, encoding, message
}
