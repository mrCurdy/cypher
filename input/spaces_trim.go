package input

import "fmt"

func TrimSpaces(message string) string {
	res := ""
	if len(message) == 0 || message == " " {
		fmt.Println("Your message is empty. I'm TrimSpaces")
		return ""
	}
	for message[0] == ' ' {
		message = message[1:]
	}
	for message[len(message)-1] == ' ' {
		message = message[:len(message)-1]
	}
	res = message
	return res
}
