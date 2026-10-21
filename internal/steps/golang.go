package steps

import (
	"fmt"

	"github.com/sikzyo/4k1/internal/message"
	"github.com/sikzyo/4k1/internal/shell"
)

func golang() {
	fmt.Println("✦ Go ✦")

	err := validateGolang()
	if err == nil {
		return
	}

	installGolang()

	disableTelemetry()

	err = validateGolang()
	if err != nil {
		message.ErrorMessage("Go no se instalo correctamente", err)
	}
}

func validateGolang() error {
	err := shell.RunCommand(false, "brew", "list", "go")
	if err != nil {
		return err
	}
	return nil
}

func installGolang() {
	err := shell.RunCommand(true, "brew", "install", "go")
	if err != nil {
		message.ErrorMessage("Error al instalar Go", err)
	}
}

func disableTelemetry() {
	err := shell.RunCommand(false, "go", "telemetry", "off")
	if err != nil {
		message.ErrorMessage("Error al deshabilitar la telemetría", err)
	}
}
