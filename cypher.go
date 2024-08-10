package main

import (
	"cypher/input"
	"cypher/reverse"
	"cypher/rot13"
	"cypher/thirdcypher"
	"fmt"
)

func cypher() {

	toEncrypt, encoding, message := input.GetInputV2()

	message = input.TrimSpaces(message)

	if toEncrypt {

		if encoding == "1" {
			fmt.Println("\nEncrypted message using Rot13:")
			fmt.Println(rot13.Encrypt_rot13(message))

		} else if encoding == "2" {

			fmt.Println("\nEncrypted message using Reverse:")
			fmt.Println(reverse.Encrypt_reverse(message))
		} else {

			fmt.Println("\nEncrypted message using Third custom Cypher:")
			fmt.Println(thirdcypher.Encrypt_thirdcypher(message))
		}

	} else {

		if encoding == "1" {

			fmt.Println("\nDecrypted message using Rot13:")
			fmt.Println(rot13.Decrypt_rot13(message))

		} else if encoding == "2" {

			fmt.Println("\nDecrypted message using Reverse:")
			fmt.Println(reverse.Decrypt_reverse(message))

		} else {

			fmt.Println("\nDecrypted message using Reverse and shift:")
			fmt.Println(thirdcypher.Decrypt_thirdcypher(message))
		}

	}

}
