package steps

import (
	"github.com/sikzyo/4k1/internal/message"
	"github.com/sikzyo/4k1/internal/shell"
)

var cliTools = []string{"fzf"}

func cli() {
	for _, tool := range cliTools {
		err := validateCli(tool)
		if err == nil {
			continue
		}

		errorMessage := "No se puedo instalar" + tool
		installCli(tool, errorMessage)

		err = validateCli(tool)
		if err != nil {
			message.ErrorMessage(errorMessage, err)
		}
	}
}

func validateCli(currentTool string) error {
	err := shell.RunCommand(false, "brew", "install", currentTool)
	if err != nil {
		return err
	}
	return nil
}

func installCli(currentTool string, errorMessage string) {
	err := shell.RunCommand(false, "brew", "install", currentTool)
	if err != nil {
		message.ErrorMessage(errorMessage, err)
	}
}
