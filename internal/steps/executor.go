package steps

import (
	"github.com/sikzyo/4k1/internal/message"
	"github.com/sikzyo/4k1/internal/shell"
)

func FullInstall() {
	for _, step := range AllSteps {
		err := shell.RunCommand(true, "clear")
		if err != nil {
			message.ErrorMessage("Error al limpiar la pantalla", err)
		}
		errorMessage, err := step.execute()
		if err != nil {
			message.ErrorMessage(errorMessage, err)
		}
	}
}
