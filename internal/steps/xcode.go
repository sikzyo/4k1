package steps

import (
	"fmt"
	"time"

	"github.com/sikzyo/4k1/internal/message"
	"github.com/sikzyo/4k1/internal/shell"
)

func Xcode() {
	fmt.Println("✦ Xcode ✦")
	err := validateXcode()
	if err == nil {
		return
	}

	installXcode()

	err = validateXcode()
	if err != nil {
		message.ErrorMessage("Xcode no se instalo correctamente", err)
	}
}

func validateXcode() error {
	err := shell.RunCommand(false, "xcode-select", "-p")
	if err != nil {
		return err
	}
	return nil
}

func installXcode() {
	fmt.Println("-> Se abrió una ventana para realizar la instalación de Xcode")
	fmt.Println("-> Realice la instalación mediante esta ventana")
	fmt.Println("-> Una vez terminado puede regresar a esta ventana de la terminal")
	err := shell.RunCommand(false, "xcode-select", "--install")
	if err != nil {
		message.ErrorMessage("Error al ejecutar comando de instalación", err)
	}
	for {
		time.Sleep(3 * time.Second)
		cmd := shell.RunCommand(false, "pgrep", "-x", "Install Command Line Developer Tools")
		if cmd != nil {
			return
		}
	}
}
