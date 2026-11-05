package steps

import (
	"fmt"

	"github.com/sikzyo/4k1/internal/message"
	"github.com/sikzyo/4k1/internal/shell"
)

var guiApps = []string{"helium-browser", "keepassxc"}

func gui() {
	fmt.Println("✦ GUI ✦")

	for _, app := range guiApps {
		err := validateGUI(app)
		if err == nil {
			continue
		}

		errorMessage := "No se puedo instalar" + app
		installGUI(app, errorMessage)

		err = validateGUI(app)
		if err != nil {
			message.ErrorMessage(errorMessage, err)
		}
	}
}

func validateGUI(currentApp string) error {
	err := shell.RunCommand(false, "brew", "install", currentApp)
	if err != nil {
		return err
	}
	return nil
}

func installGUI(currentApp string, errorMessage string) {
	err := shell.RunCommand(true, "brew", "install", "--cask", currentApp)
	if err != nil {
		message.ErrorMessage(errorMessage, err)
	}
}
