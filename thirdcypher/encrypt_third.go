package thirdcypher

import (
	"cyphertool/rot13"
)

func Encrypt_thirdcypher(s string) string {
	var result string
	for _, v := range s {
		result = string(v) + result
	}
	result = rot13.Decrypt_rot13(result)
	return result
}
