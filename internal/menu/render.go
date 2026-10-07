package menu

import (
	"fmt"

	"github.com/sikzyo/4k1/internal/message"
	"github.com/sikzyo/4k1/internal/shell"
)

func ShowMenu(currentMenu Menu) {
	err := shell.RunCommand(true, "clear")
	if err != nil {
		message.ErrorMessage("Error al limpiar la pantalla", err)
	}
	if currentMenu.Logo {
		showLogo()
	}
	showDivider()
	fmt.Println(currentMenu.Title)
	showDivider()
	showOptions(currentMenu.Options)
	showDivider()
	fmt.Println("0", "-", currentMenu.Fallback)
	showDivider()
}

func showLogo() {
	fmt.Println("   __ __  __  ___")
	fmt.Println("  / // / / /_<  /")
	fmt.Println(" / // /_/ //_/ / ")
	fmt.Println("/__  __/ ,< / /  ")
	fmt.Println("  /_/ /_/|_/_/   ")
}

func showDivider() {
	fmt.Println("--------------------")
}

func showOptions(options []MenuOptions) {
	for index, option := range options {
		fmt.Println(index+1, "-", option.Title)
	}
}
