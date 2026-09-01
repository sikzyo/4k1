package tasks

import (
	"fmt"
	"os"

	"github.com/sikzyo/4k1/internal/data"
	"github.com/sikzyo/4k1/internal/execute"
)

func FullInstall() {

	for _, step := range data.AllSteps {
		execute.Command("clear")
		cmd := step.Step()

		if cmd != nil {
			fmt.Println()
			fmt.Println("-> Ocurrió un error al momento de ejecutar un proceso, por seguridad el script se va a detener")
			os.Exit(1)
		}
	}
}
