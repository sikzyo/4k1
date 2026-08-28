package main

import (
	"github.com/sikzyo/4k1/internal/menu"
	"github.com/sikzyo/4k1/internal/models"
)

func main() {

	MainMenu := models.MenuModel{
		ShowLogo: true,
		Title:    "Menu Principal",
	}

	menu.ShowMenu(MainMenu)
}
