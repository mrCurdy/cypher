package rot13

func Encrypt_rot13(s string) string {
	// need to add filters for numbers, spaces, capitals ...
	res := ""

	for _, char := range s {

		// checking if lowcase char
		if char >= 'a' && char <= 'z' {
			res += shiftBy5(char, 'z')

			// checking if capcase char
		} else if char >= 'A' && char <= 'Z' {
			res += shiftBy5(char, 'Z')

			// all other symbols including numbers and spaces
		} else {
			res += string(char)
			continue
		}
	}

	return res
}

// function for shifting char position in alphabet
func shiftBy5(char rune, end rune) string {

	letter := int32(5) + char
	for letter > end {
		letter -= 26
	}

	return string(letter)
}
