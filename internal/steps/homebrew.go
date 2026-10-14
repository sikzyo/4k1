package steps

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/sikzyo/4k1/internal/download"
	"github.com/sikzyo/4k1/internal/shell"
)

func homebrew() (errorMessage string, err error) {
	fmt.Println("✦ Homebrew ✦")
	err = validateHomebrew()
	if err == nil {
		return "Homebrew ya se encuentra instalado", nil
	}

	err = installHomebrew()
	if err != nil {
		return "Error al momento de instalar Homebrew", err
	}

	err = configHomebrew()
	if err != nil {
		return "Error al configurar Homebrew", err
	}

	err = validateHomebrew()
	if err != nil {
		return "Homebrew no se puedo instalar de manera correcta", err
	}

	return "", nil
}

func validateHomebrew() error {
	err := shell.RunCommand(false, "brew", "--version")
	if err != nil {
		return err
	}
	return nil
}

func installHomebrew() error {
	err := shell.RunCommand(false, "/bin/bash", "-c", "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)")
	if err != nil {
		return err
	}
	return nil
}

func configHomebrew() error {
	userDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	path := filepath.Join(userDir, "test/prueba.md")

	err = download.DownloadFiles("https://raw.githubusercontent.com/sikzyo/4k1/heads/main/config/.zprofile", path)
	if err != nil {
		return err
	}
	return nil
}
