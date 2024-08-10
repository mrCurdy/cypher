package reverse

func Decrypt_reverse(s string) string {
	reversedRunes := make([]rune, len(s))
	
	for i, r := range s {
		reversedRunes[len(s)-1-i] = ReverseAlphabet2(r)
	}
	
	return string(reversedRunes)
	
}
func ReverseAlphabet2(ch rune) rune {
	var reverse rune
	if ch >= 'a' && ch <= 'z' {
		reverse = 'z' - (ch - 'a')
	} else if ch >= 'A' && ch <= 'Z' {
		reverse = 'Z' - (ch - 'A')
	} else {
		return ch
	}
	
	return reverse 
}
