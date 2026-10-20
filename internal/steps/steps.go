// Package steps is to perform the installation steps in order
package steps

import (
	"os"

	"github.com/sikzyo/4k1/internal/message"
)

type step struct {
	name    string
	execute func()
}

var (
	gitEmail string
	userDir  string
)

var AllSteps = []step{
	{name: "Xcode", execute: Xcode},
	{name: "Homebrew", execute: homebrew},
	{name: "Git", execute: git},
	{name: "SSH", execute: ssh},
	{name: "Fonts", execute: fonts},
	{name: "Node", execute: node},
	{name: "Go", execute: golang},
}

func setUserDir() {
	var err error

	userDir, err = os.UserHomeDir()
	if err != nil {
		message.ErrorMessage("Error al obtener carpeta del usuario", err)
	}
}
