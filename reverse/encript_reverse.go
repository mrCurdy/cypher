package reverse

func Encrypt_reverse(message string) string {
	res := ""

	for _, char := range message {
		char = ReverseAlphabet(char)
		res += string(char)
	}
	return res
}

//function was replaced. Original worked with bugs
// reversedRunes := make([]rune, len(s))

// for i, r := range s {
// 	reversedRunes[len(s)-1-i] = ReverseAlphabet(r)
// }
