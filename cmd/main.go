package main

import (
	"github.com/sikzyo/4k1/internal/menu"
)

func main() {
	mainMenu := menu.Menu{
		Name:        "Menu principal",
		ShowLogo:    true,
		MenuOptions: []string{"Instalación completa", "Instalación por partes"},
	}

	menu.ShowMenu(mainMenu)
}
