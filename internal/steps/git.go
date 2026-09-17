package steps

import (
	"errors"
	"fmt"

	"github.com/sikzyo/4k1/internal/execute"
	"github.com/sikzyo/4k1/internal/input"
)

var user_git string
var email_git string

func Git() error {
	fmt.Println("✦ Git ✦")

	err := installGit()
	if err != nil {
		return err
	}

	config, err := gitMenu()
	if err != nil {
		return err
	}

	if !config {
		return nil
	}

	return gitConfig()
}

func installGit() error {
	fmt.Println("-> Validando instalación de git")

	err := execute.Command("brew", "list", "git")

	if err == nil {
		fmt.Println("-> Git ya se encuentra instalado")
		return nil
	}

	fmt.Println("-> Instalando Git mediante Homebrew")
	err = execute.CommandNull("brew", "install", "git")

	if err != nil {
		return errors.New("La instalación de Git fallo")
	}

	fmt.Println("-> Git instalado correctamente")

	return nil
}

func gitMenu() (bool, error) {
	for {
		fmt.Println("-> ¿Deseas configurar git? [y/n]")
		response, err := input.GetInput()

		if err != nil {
			return false, err
		}

		switch response {
		case "y", "Y":
			return true, nil
		case "n", "N":
			return false, nil
		default:
			fmt.Println("-> Opción no valida, por favor inténtalo nuevamente")
		}
	}
}

func gitConfig() error {
	fmt.Println("✦ Configuración de Git ✦")

	fmt.Println("-> Ingresa tu nombre de usuario para Git")
	user_git, err := input.GetInput()
	if err != nil {
		return errors.New("Error al registrar el nombre")
	}

	fmt.Println("-> Ingresa tu correo para Git")
	email_git, err := input.GetInput()
	if err != nil {
		return errors.New("Error al registrar el correo")
	}

	fmt.Println("-> Aplicando configuraciones de Git")
	err = execute.Command("git", "config", "--global", "user.name", user_git)
	if err != nil {
		return errors.New("Error al configurar el nombre")
	}

	err = execute.Command("git", "config", "--global", "user.email", email_git)
	if err != nil {
		return errors.New("Error al configurar el correo")
	}

	err = execute.Command("git", "config", "--global", "init.defaultBranch", "main")
	if err != nil {
		return errors.New("Error al configurar rama por defecto")
	}

	fmt.Println("✦ Git se configuro correctamente")
	return nil
}
