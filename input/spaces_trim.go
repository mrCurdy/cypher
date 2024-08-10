package input

func TrimSpaces(message string) string {
	res := ""

	for message[0] == ' ' {
		message = message[1:]
	}
	for message[len(message)-1] == ' ' {
		message = message[:len(message)-1]
	}
	res = message
	return res
}
