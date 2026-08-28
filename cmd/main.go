package main

import (
	"github.com/sikzyo/4k1/internal/menu"
	"github.com/sikzyo/4k1/internal/models"
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
	}

	menu.ShowMenu(MainMenu)
}
