package main

import (
	"github.com/sikzyo/4k1/internal/menu"
)

func main() {
	mainMenu := menu.Menu{
		Title: "✦ Menu principal ✦",
		Logo:  true,
		Options: []menu.MenuOptions{
			{Title: "Instalación completa"},
			{Title: "Instalación por secciones"},
		},
		Fallback: "Salir de la aplicación",
	}

	menu.ShowMenu(mainMenu)
}
