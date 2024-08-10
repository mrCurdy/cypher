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
			fmt.Println("Encrypted message using Rot13:")
			fmt.Println(rot13.Encrypt_rot13(message))

		} else if encoding == "2" {

			fmt.Println("Encrypted message using Reverse:")
			fmt.Println(reverse.Encrypt_reverse(message))
		} else {

			fmt.Println("Encrypted message using Third custom Cypher:")
			fmt.Println(thirdcypher.Encrypt_thirdcypher(message))
		}

	} else {

		if encoding == "1" {

			fmt.Println("Decrypted message using Rot13:")
			fmt.Println(rot13.Decrypt_rot13(message))

		} else if encoding == "2" {

			fmt.Println("Decrypted message using Reverse:")
			fmt.Println(reverse.Decrypt_reverse(message))

		} else {

			fmt.Println("Decrypted message using Third custom Cypher:")
			fmt.Println(thirdcypher.Decrypt_thirdcypher(message))
		}

	}

}
