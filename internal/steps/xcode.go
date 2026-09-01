package steps

import (
	"fmt"

	"github.com/sikzyo/4k1/internal/execute"
)

func InstallXcode() error {
	fmt.Println("✦ Xcode ✦")

	fmt.Println("-> Validando versión de xcode")
	err := execute.Command("xcode-select", "-v")

	if err == nil {
		fmt.Println("-> Xcode instalado correctamente")
		return nil
	}

	fmt.Println("-> Instalando xcode")
	err = execute.Command("xcode-select", "--install")

	return err
}
