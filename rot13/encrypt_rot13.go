package rot13

func Encrypt_rot13(s string) string {

	res := ""

	for _, char := range s {

		// checking if lowcase char
		if char >= 'a' && char <= 'z' {
			res += shiftBy13(char, 'z')

			// checking if capcase char
		} else if char >= 'A' && char <= 'Z' {
			res += shiftBy13(char, 'Z')

			// all other symbols including numbers and spaces
		} else {
			res += string(char)
			continue
		}
	}

	return res
}

// function for shifting char position in alphabet
func shiftBy13(char rune, end rune) string {

	letter := int32(13) + char
	for letter > end {
		letter -= 26
	}

	return string(letter)
}
