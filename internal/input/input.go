// Package input is used to manage user input
package input

import (
	"bufio"
	"io"
	"os"
)

var scanner = bufio.NewScanner(os.Stdin)

func GetInput() (string, error) {
	if scanner.Scan() {
		return scanner.Text(), nil
	}

	err := scanner.Err()
	if err != nil {
		return "", err
	}

	return "", io.EOF
}
