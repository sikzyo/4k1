package steps

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/sikzyo/4k1/internal/download"
	"github.com/sikzyo/4k1/internal/message"
	"github.com/sikzyo/4k1/internal/shell"
)

func homebrew() {
	fmt.Println("✦ Homebrew ✦")
	err := validateHomebrew()
	if err == nil {
		return
	}

	installHomebrew()

	configHomebrew()

	err = validateHomebrew()
	if err != nil {
		message.ErrorMessage("No se instalo correctamente Homebrew", err)
	}
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
		message.ErrorMessage("El comando de instalación presento un fallo", err)
	}
	return nil
}

func configHomebrew() {
	userDir, err := os.UserHomeDir()
	if err != nil {
		message.ErrorMessage("No se pudo obtener la carpeta del usuario", err)
	}
	path := filepath.Join(userDir, "test/prueba.md")

	err = download.DownloadFiles("https://raw.githubusercontent.com/sikzyo/4k1/heads/main/config/.zprofile", path)
	if err != nil {
		message.ErrorMessage("La descarga del archivo de configuración fallo", err)
	}
}
