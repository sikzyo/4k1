package menu

import (
	"fmt"

	"github.com/sikzyo/4k1/internal/models"
)

func ShowMenu(Menu models.MenuModel) {
	if Menu.ShowLogo {
		ShowLogo()
	}
	ShowDivider()
	fmt.Println(Menu.Title)
	ShowDivider()
	for indice, option := range Menu.Options {
		fmt.Print(indice, ") ", option, "\n")
	}
	ShowDivider()
	fmt.Println("0)", Menu.Exit)
	ShowDivider()
}
