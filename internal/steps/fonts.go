package steps

import (
	"fmt"

	"github.com/sikzyo/4k1/internal/message"
	"github.com/sikzyo/4k1/internal/shell"
)

var currentFont string = "font-jetbrains-mono-nerd-font"

func fonts() {
	fmt.Println("✦ Fuentes")

	err := validateFont()
	if err == nil {
		return
	}

	installFont()

	err = validateFont()
	if err != nil {
		message.ErrorMessage("No se instalo no fue correcta", err)
	}
}

func validateFont() error {
	err := shell.RunCommand(false, "brew", "list", currentFont)
	if err != nil {
		return err
	}
	return nil
}

func installFont() {
	err := shell.RunCommand(false, "brew", "install", "--cask", currentFont)
	if err != nil {
		message.ErrorMessage("No se pudo instalar la fuente", err)
	}
}
