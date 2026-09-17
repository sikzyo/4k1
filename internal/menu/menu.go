package menu

import (
	"fmt"
	"time"

	"github.com/sikzyo/4k1/internal/execute"
	"github.com/sikzyo/4k1/internal/models"
)

func ShowMenu(Menu models.MenuModel) {
	execute.Command("clear")
	if Menu.ShowLogo {
		showLogo()
	}
	showDivider()
	fmt.Println(Menu.Title)
	showDivider()
	for indice, option := range Menu.Options {
		fmt.Print(indice+1, ") ", option, "\n")
	}
	showDivider()
	fmt.Println("0)", Menu.Exit)
	showDivider()
}

func ShowMessages(message string, delay int) {
	execute.Command("clear")
	showDivider()
	fmt.Println("△ Mensaje:", message, "△")
	showDivider()
	time.Sleep(time.Duration(delay) * time.Second)
}

func showDivider() {
	fmt.Println("——————————————")
}

func showLogo() {
	fmt.Println("   __ __  __  ___")
	fmt.Println("  / // / / /_<  /")
	fmt.Println(" / // /_/ //_/ / ")
	fmt.Println("/__  __/ ,< / /  ")
	fmt.Println("  /_/ /_/|_/_/   ")
}
