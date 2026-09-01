package main

import (
	"fmt"

	"github.com/sikzyo/4k1/internal/input"
	"github.com/sikzyo/4k1/internal/menu"
	"github.com/sikzyo/4k1/internal/models"
	"github.com/sikzyo/4k1/internal/tasks"
)

func main() {

	MainMenuOptions := []string{
		"Instalación completa",
		"Instalación por módulos",
	}

	MainMenu := models.MenuModel{
		ShowLogo: true,
		Title:    "Menu Principal",
		Options:  MainMenuOptions,
		Exit:     "Salir de 4k1",
	}

	close := false
	for !close {
		menu.ShowMenu(MainMenu)
		response, err := input.GetInput()

		if err != nil {
			fmt.Println("-> Error:", err)

		}

		switch response {
		case 1:
			tasks.FullInstall()
			close = true
		case 0:
			fmt.Println("Gracias por usar 4k1")
			close = true
		default:
			fmt.Println("Opción no valida, intenta nuevamente")
		}
	}

}
