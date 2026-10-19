package steps

import (
	"fmt"
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

func installHomebrew() {
	err := shell.RunCommand(false, "/bin/bash", "-c", "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)")
	if err != nil {
		message.ErrorMessage("El comando de instalación presento un fallo", err)
	}
}

func configHomebrew() {
	if userDir == "" {
		setUserDir()
	}

	path := filepath.Join(userDir, ".zprofile")
	err := download.DownloadFiles("https://raw.githubusercontent.com/sikzyo/4k1/heads/main/config/.zprofile", path)
	if err != nil {
		message.ErrorMessage("La descarga del archivo de configuración fallo", err)
	}
}
