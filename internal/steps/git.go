package steps

import (
	"fmt"

	"github.com/sikzyo/4k1/internal/input"
	"github.com/sikzyo/4k1/internal/message"
	"github.com/sikzyo/4k1/internal/shell"
)

func git() {
	fmt.Println("✦ Git ✦")
	err := validateGit()
	if err == nil {
		return
	}

	installGit()

	err = validateGit()
	if err != nil {
		message.ErrorMessage("No se instalo correctamente Git mediante Homebrew", err)
	}

	selectOption()
}

func validateGit() error {
	err := shell.RunCommand(false, "brew", "list", "git")
	if err != nil {
		return err
	}
	return nil
}

func installGit() {
	err := shell.RunCommand(true, "brew", "install", "git")
	if err != nil {
		message.ErrorMessage("El comando para instalar Git fallo", err)
	}
}

func selectOption() {
	for {
		fmt.Println("✦ ¿Desear configurar Git? [y/n]")
		fmt.Print("-> ")
		response := input.GetInput()
		switch response {
		case "y", "Y":
			configGit()
			return
		case "n", "N":
			return
		default:
			message.TimeMessage("La opción seleccionada no esta disponible", 3)
		}
	}
}

func configGit() {
	setGitUser()
	setGitEmail()
	setDefaultBranch()
}

func setGitEmail() {
	fmt.Println("✦ Ingresa tu correo electrónico")
	fmt.Print("-> ")

	gitEmail = input.GetInput()
	err := shell.RunCommand(false, "git", "config", "--global", "user.email", gitEmail)
	if err != nil {
		message.ErrorMessage("Error al configurar el correo electrónico", err)
	}
}

func setGitUser() {
	fmt.Println("✦ Ingresa tu nombre de usuario")
	fmt.Print("-> ")

	gitUsername := input.GetInput()
	err := shell.RunCommand(false, "git", "config", "--global", "user.name", gitUsername)
	if err != nil {
		message.ErrorMessage("Error al configurar nombre de usuario", err)
	}
}

func setDefaultBranch() {
	err := shell.RunCommand(false, "git", "config", "--global", "init.defaultBranch", "main")
	if err != nil {
		message.ErrorMessage("Error al configurar rama principal", err)
	}
}
