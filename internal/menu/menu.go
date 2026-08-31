package menu

import (
	"fmt"
	"os"

	"github.com/sikzyo/4k1/internal/execute"
	"github.com/sikzyo/4k1/internal/models"
)

func ShowMenu(Menu models.MenuModel) {
	err := execute.Command("clear")
	if err != nil {
		fmt.Println("Un error ocurrió al iniciar el programa")
		os.Exit(1)
	}
	if Menu.ShowLogo {
		ShowLogo()
	}
	ShowDivider()
	fmt.Println(Menu.Title)
	ShowDivider()
	for indice, option := range Menu.Options {
		fmt.Print(indice+1, ") ", option, "\n")
	}
	ShowDivider()
	fmt.Println("0)", Menu.Exit)
	ShowDivider()
}
