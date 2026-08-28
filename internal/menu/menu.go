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
}
