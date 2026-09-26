package menu

import "fmt"

func ShowMenu(currentMenu Menu) {
	if currentMenu.Logo {
		showLogo()
	}
	showDivider()
	fmt.Println(currentMenu.Title)
	showDivider()
	showOptions(currentMenu.Options)
	showDivider()
	fmt.Println("0", "-", currentMenu.Fallback)
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
