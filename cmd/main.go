package main

import (
	"fmt"
	"os"

	"github.com/sikzyo/4k1/internal/input"
	"github.com/sikzyo/4k1/internal/menu"
	"github.com/sikzyo/4k1/internal/models"
	"github.com/sikzyo/4k1/internal/tasks"
)

func main() {

	MainMenuOptions := []string{
		"Instalación completa",
		// "Instalación por módulos",
	}

	MainMenu := models.MenuModel{
		ShowLogo: true,
		Title:    "Menu Principal",
		Options:  MainMenuOptions,
		Exit:     "Salir de 4k1",
	}

	menuOptions(MainMenu)
}

func menuOptions(mainMenu models.MenuModel) {
	for {
		menu.ShowMenu(mainMenu)
		responde, err := input.GetInput()

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		switch responde {
		case "1":
			tasks.FullInstall()
		case "0":
			fmt.Println("-> Gracias por utilizar 4k1")
			return
		default:
			menu.ShowMessages("La opción que seleccionaste no es correcta, por favor inténtalo nuevamente", 4)
		}

	}
}
