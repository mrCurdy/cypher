package input

func TrimSpaces(message string) string {
	res := ""
	// if len(message) == 0 {
	// 	fmt.Println("empty")
	// 	return ""
	// }
	for message[0] == ' ' {
		message = message[1:]
	}
	for message[len(message)] == ' ' {
		message = message[:len(message)-1]
	}
	res = message
	return res
}
