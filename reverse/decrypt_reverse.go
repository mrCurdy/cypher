package reverse

func Decrypt_reverse(s string) string {
	reversedRunes := make([]rune, len(s))

	for i, r := range s {
		reversedRunes[len(s)-1-i] = ReverseAlphabet(r)
	}

	return string(reversedRunes)

}
