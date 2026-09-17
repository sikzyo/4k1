package steps

import (
	"errors"
	"fmt"
	"time"

	"github.com/sikzyo/4k1/internal/execute"
)

func Xcode() error {
	fmt.Println("✦ Xcode ✦")

	err := validateXcode()

	if err == nil {
		return nil
	}

	err = installXcode()

	if err != nil {
		return errors.New("No se puede iniciar la instalación de Xcode")
	}

	err = validateXcode()

	if err != nil {
		return errors.New("La instalación de Xcode presento un problema")
	}

	fmt.Println("✦ La instalación de Xcode se ejecuto de manera correcta")

	return nil
}

func validateXcode() error {
	fmt.Println("-> Validando instalación de xcode")
	err := execute.Command("xcode-select", "-p")

	if err != nil {
		return err
	}
	return nil
}

func installXcode() error {
	fmt.Println("-> Instalando xcode")
	err := execute.Command("xcode-select", "--install")

	if err != nil {
		return err
	}

	fmt.Println("✦ Se abrió una ventana para la instalación de Xcode")
	fmt.Println("✦ Continua en esa ventana el proceso de instalación")
	fmt.Println("✦ Al terminar la instalación de Xcode, regresa para continuar con la configuración")

	for true {
		// Comando para validar si el proceso de instalación termino
		err = execute.CommandNull("pgrep", "-f", "Install Command Line Developer Tools")
		if err != nil {
			break
		}

		time.Sleep(5 * time.Second)
	}
	return nil
}
