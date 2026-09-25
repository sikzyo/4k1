// Package menu utilities to render the menu
package menu

import "fmt"

func ShowMenu() {
	showIcon()
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
