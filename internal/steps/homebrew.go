package steps

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/sikzyo/4k1/internal/dir"
	"github.com/sikzyo/4k1/internal/download"
	"github.com/sikzyo/4k1/internal/message"
	"github.com/sikzyo/4k1/internal/shell"
)

const homebrewInstallURL = "https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh"

func homebrew() {
	fmt.Println("✦ Homebrew ✦")
	err := validateHomebrew()
	if err == nil {
		return
	}

	installHomebrew()

	configHomebrew()

	updatePath()

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
	tempDir, err := os.MkdirTemp("", "homebrew-installation-*")
	if err != nil {
		message.ErrorMessage("No se pudo crear la carpeta temporal", err)
	}

	scriptPath := filepath.Join(tempDir, "install.sh")

	err = download.DownloadFiles(homebrewInstallURL, scriptPath)
	if err != nil {
		removeTempDir(tempDir)
		message.ErrorMessage("La descarga del archivo de instalación fallo", err)
	}

	err = shell.RunCommand(true, "/bin/bash", scriptPath)

	removeTempDir(tempDir)

	if err != nil {
		message.ErrorMessage("El comando de instalación presento un fallo", err)
	}
}

func removeTempDir(tempDir string) {
	err := os.RemoveAll(tempDir)
	if err != nil {
		message.ErrorMessage("No se pudo eliminar la carpeta temporal", err)
	}
}

func updatePath() {
	err := os.Setenv("PATH", "/opt/homebrew/bin:/opt/homebrew/sbin:"+os.Getenv("PATH"))
	if err != nil {
		message.ErrorMessage("Error al actualizar el path", err)
	}
}

func configHomebrew() {
	if dir.UserDir == "" {
		dir.SetUserDir()
	}

	path := filepath.Join(dir.UserDir, ".zprofile")
	err := download.DownloadFiles("https://raw.githubusercontent.com/sikzyo/4k1/heads/main/config/.zprofile", path)
	if err != nil {
		message.ErrorMessage("La descarga del archivo de configuración fallo", err)
	}
}
