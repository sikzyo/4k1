package main

import (
	"fmt"

	"github.com/sikzyo/4k1/internal/input"
	"github.com/sikzyo/4k1/internal/menu"
	"github.com/sikzyo/4k1/internal/message"
	"github.com/sikzyo/4k1/internal/steps"
)

var mainMenu = menu.Menu{
	Title: "✦ Menu principal ✦",
	Logo:  true,
	Options: []menu.MenuOptions{
		{Title: "Instalación completa"},
	},
	Fallback: "Salir",
}
var status string

func main() {
	run()
}

func run() {
	for {
		menu.ShowMenu(mainMenu)

		if status != "" {
			fmt.Println(status)
			status = ""
		}

		fmt.Println("✦ Selecciona una opción del menu")
		fmt.Print("-> ")

		response, err := input.GetInput()
		if err != nil {
			message.ErrorMessage("No se pudo obtener el input correctamente", err)
		}

		switch response {
		case "1":
			steps.FullInstall()
		case "0":
			fmt.Println("✦ Gracias por utilizar 4k1")
			return
		default:
			status = "La opción no se encuentra disponible"
		}
	}
}
