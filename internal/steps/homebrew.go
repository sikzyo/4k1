package steps

import (
	"errors"
	"fmt"
	"os"

	"github.com/sikzyo/4k1/internal/execute"
)

func InstallHomebrew() error {
	fmt.Println("✦ Homebrew ✦")

	fmt.Println("-> Validando Homebrew")
	err := execute.Command("command", "-v", "homebrew")

	if err != nil {
		fmt.Println("-> Homebrew ya se encuentra instalado")
		return nil
	}

	fmt.Println("-> Instalando Homebrew")

	err = execute.Command("/bin/bash", "-c", "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)")

	if err != nil {
		return errors.New("El script de instalación de Homebrew fallo")
	}

	fmt.Println("-> Creando archivo .zprofile")

	file, err := os.Create("$HOME/.zprofile")

	if err != nil {
		return errors.New("Error al crearr el archivo .zprofile")
	}

	defer file.Close()

	_, err = file.WriteString("eval '$(/opt/homebrew/bin/brew shellenv)'")

	if err != nil {
		return errors.New("Error al momento de configurar .zprofile")
	}

	fmt.Println("-> Validando instalación de Homebrew")
	err = execute.Command("brew", "-v")

	if err != nil {
		return errors.New("No se pudo validar correctamente la instalación de Homebrew")
	}

	fmt.Println("✦ Homebrew se instalo correctamente")

	return nil
}
