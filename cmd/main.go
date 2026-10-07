package main

import (
	"fmt"
	"os"

	"github.com/sikzyo/4k1/internal/input"
	"github.com/sikzyo/4k1/internal/menu"
	"github.com/sikzyo/4k1/internal/message"
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

	for {
		menu.ShowMenu(mainMenu)
		selectOption()
	}
}

func selectOption() {
	fmt.Println("✦ Selecciona una opción del menu")
	fmt.Print("-> ")
	response := input.GetInput()
	switch response {
	case "1":
		fmt.Print("Opción aun no implementada")
	case "2":
		fmt.Print("Algún día implementare esto")
	case "0":
		fmt.Println("✦ Gracias por usar 4k1")
		os.Exit(0)
	default:
		message.TimeMessage("La opción seleccionada no esta disponible", 3)
	}
}
