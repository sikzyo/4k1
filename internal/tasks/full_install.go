package tasks

import (
	"os"

	"github.com/sikzyo/4k1/internal/data"
	"github.com/sikzyo/4k1/internal/execute"
	"github.com/sikzyo/4k1/internal/menu"
)

func FullInstall() {

	for _, step := range data.AllSteps {
		execute.Command("clear")
		err := step.Step()

		if err != nil {
			menu.ShowMessages(err.Error(), 0)
			os.Exit(1)
		}
	}
}
