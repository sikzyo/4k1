// Package menu utilities to render the menu
package menu

import "fmt"

type Menu struct {
	Name        string
	ShowLogo    bool
	MenuOptions []string
	Fallback    func()
}

func ShowMenu(currentMenu Menu) {
	if currentMenu.ShowLogo {
		showIcon()
	}
	fmt.Println(currentMenu.Name)
}

func showIcon() {
	fmt.Println("   __ __  __  ___")
	fmt.Println("  / // / / /_<  /")
	fmt.Println(" / // /_/ //_/ / ")
	fmt.Println("/__  __/ ,< / /  ")
	fmt.Println("  /_/ /_/|_/_/   ")
}

// TODO: Se debe poder pasar las opciones al menu y este las debe de mostrar
func showOptions() {
}
