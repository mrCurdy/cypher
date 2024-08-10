package rot13

func Decrypt_rot13(s string) string {
	// need to add filters for numbers, spaces, capitals ...
	res := ""

	for _, char := range s {

		// checking if lowcase char
		if char >= 'a' && char <= 'z' {
			res += shiftBackBy13(char, 'a')

			// checking if capcase char
		} else if char >= 'A' && char <= 'Z' {
			res += shiftBackBy13(char, 'A')

			// all other symbols including numbers and spaces
		} else {
			res += string(char)
			continue
		}
	}

	return res
}

// function for shifting char position in alphabet
func shiftBackBy13(char rune, end rune) string {

	letter := char - int32(13)
	for letter < end {
		letter += 26
	}

	return string(letter)
}
