// Package input is used to manage user input
package input

import (
	"bufio"
	"os"
)

func GetInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
