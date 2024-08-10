package main

import (
	"cypher/input"
	"cypher/reverse"
	"cypher/rot13"
	"cypher/thirdcypher"
	"fmt"
)

func main() {
	toEncrypt, encoding, message := input.getInput()
	message = input.TrimSpaces(message)
	if toEncrypt {
		if encoding == "1" {
			fmt.Println("Encrypted message using Rot13:")
			fmt.Print(rot13.Encrypt_rot13(message))
		} else if encoding == "2" {
			fmt.Println("Encrypted message using Reverse:")
			fmt.Print(reverse.Encrypt_reverse(message))
		} else {
			fmt.Println("Encrypted message using Third custom Cypher:")
			fmt.Print(thirdcypher.Encrypt_thirdcypher(message))
		}
	} else {
		if encoding == "1" {
			fmt.Println("Decrypted message using Rot13:")
			fmt.Print(rot13.Decrypt_rot13(message))
		} else if encoding == "2" {
			fmt.Println("Decrypted message using Reverse:")
			fmt.Print(reverse.Decrypt_reverse(message))
		} else {
			fmt.Println("Decrypted message using Third custom Cypher:")
			fmt.Print(thirdcypher.Decrypt_thirdcypher(message))
		}
	}

}
