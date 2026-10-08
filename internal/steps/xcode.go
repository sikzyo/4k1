package steps

import (
	"fmt"
	"time"

	"github.com/sikzyo/4k1/internal/shell"
)

func Xcode() (errorMessage string, err error) {
	err = shell.RunCommand(true, "clear")
	if err != nil {
		return "No se pudo limpiar la pantalla correctamente", err
	}

	fmt.Println("✦ Xcode ✦")
	err = validatedXcode()
	if err == nil {
		return "Xcode ya se encuentra instalado", nil
	}

	err = installXcode()
	if err != nil {
		return "Error al momento de instalar Xcode", err
	}

	err = validatedXcode()
	if err != nil {
		return "La instalación de Xcode no se pudo realizar de manera correcta", err
	}

	return "", nil
}

func validatedXcode() error {
	err := shell.RunCommand(false, "xcode-select", "-p")
	if err != nil {
		return err
	}
	return nil
}

func installXcode() error {
	fmt.Println("-> Se abrió una ventana para realizar la instalación de Xcode")
	fmt.Println("-> Realice la instalación mediante esta ventana")
	fmt.Println("-> Una vez terminado puede regresar a esta ventana de la terminal")
	err := shell.RunCommand(false, "xcode-select", "--install")
	if err != nil {
		return err
	}
	for {
		time.Sleep(3 * time.Second)
		cmd := shell.RunCommand(false, "pgrep", "-x", "Install Command Line Developer Tools")
		if cmd != nil {
			return nil
		}
	}
}
