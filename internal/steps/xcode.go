package steps

import (
	"errors"
	"fmt"
	"time"

	"github.com/sikzyo/4k1/internal/execute"
)

func InstallXcode() error {
	fmt.Println("✦ Xcode ✦")

	fmt.Println("-> Validando versión de xcode")
	err := execute.Command("xcode-select", "-p")

	if err == nil {
		fmt.Println("-> Xcode ya se encuentra instalado")
		return nil
	}

	fmt.Println("-> Instalando xcode")
	err = execute.Command("xcode-select", "--install")

	if err != nil {
		return errors.New("Error al instalar Xcode")
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

	err = execute.Command("xcode-select", "-p")

	if err != nil {
		return errors.New("La instalación de Xcode no puedo realizar correctamente")
	}

	fmt.Println("✦ La instalación de Xcode se ejecuto de manera correcta")

	return nil
}
