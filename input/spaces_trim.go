package input

func TrimSpaces(message string) string {
	//adding to message two marks
	//this is filtering messages such as "          " or "        1              " .
	message += "!!"
	//remowing spaces before message
	for message[0] == ' ' {
		message = message[1:]
	}
	//removing two marks
	res := message[:len(message)-2]

	if len(res) == 0 || res == " " {
		// fmt.Println("Your message is empty. I'm TrimSpaces")
		return ""
	}
	//removing spaces after message
	for res[len(res)-1] == ' ' {
		res = res[:len(res)-1]
	}

	return res
}

// for message[len(message)-1] == ' ' {
// 	message = message[:len(message)-1]
// }
