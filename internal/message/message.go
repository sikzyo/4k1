// Package message is for show message to user
package message

import (
	"fmt"
	"os"
	"time"

	"github.com/sikzyo/4k1/internal/shell"
)

func ErrorMessage(message string, error error) {
	err := shell.RunCommand(true, "clear")
	if err != nil {
		os.Exit(1)
	}
	fmt.Println("->", message)
	fmt.Println("->", error)
	os.Exit(1)
}

func TimeMessage(message string, timer int) {
	err := shell.RunCommand(true, "clear")
	if err != nil {
		ErrorMessage("Error al limpiar la pantalla", err)
	}

	fmt.Println("->", message)
	time.Sleep(time.Duration(timer * int(time.Second)))
}
