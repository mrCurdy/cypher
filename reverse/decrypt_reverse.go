package reverse

func Decrypt_reverse(message string) string {
	res := ""

	for _, char := range message {
		char = ReverseAlphabet(char)
		res += string(char)
	}
	return res
}

//Totally copy encript function
