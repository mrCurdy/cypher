package input

import (
	"bufio"
	"os"
)

func Scanner() string {
	input := ""
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = TrimSpaces(scanner.Text())
	return input
}
