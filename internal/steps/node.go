package steps

import (
	"fmt"

	"github.com/sikzyo/4k1/internal/message"
	"github.com/sikzyo/4k1/internal/shell"
)

var javaScriptTools = []string{"node", "pnpm"}

func node() {
	fmt.Println("✦ Node y Pnpm")

	for _, tool := range javaScriptTools {
		err := validateNode(tool)
		if err == nil {
			continue
		}

		installNode(tool)

		errorMessage := "No se pudo instalar correctamente " + tool
		err = validateNode(tool)
		if err != nil {
			message.ErrorMessage(errorMessage, err)
		}
	}
}

func validateNode(currentTool string) error {
	err := shell.RunCommand(false, "brew", "list", currentTool)
	if err != nil {
		return err
	}
	return nil
}

func installNode(currentTool string) {
	errorMessage := "Error al instalar " + currentTool
	err := shell.RunCommand(false, "brew", "install", currentTool)
	if err != nil {
		message.ErrorMessage(errorMessage, err)
	}
}
