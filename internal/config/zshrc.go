package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/sikzyo/4k1/internal/dir"
	"github.com/sikzyo/4k1/internal/download"
	"github.com/sikzyo/4k1/internal/message"
)

func zshrcConfig() {
	fmt.Println("✦ zsh ✦")

	getZshrc()
	getExports()
	getAlias()
	getTools()
}

func getTools() {
	if dir.ConfigDir == "" {
		dir.SetConfigDir()
	}

	toolsFile := filepath.Join(dir.ConfigDir, "zsh", "tools.zsh")
	downloadURL := download.GetDownloadURL("zsh", "tools.zsh")

	err := download.DownloadFiles(downloadURL, toolsFile)
	if err != nil {
		message.ErrorMessage("No se pudo descargar el archivo tools.zsh correctamente", err)
	}
}

func getExports() {
	if dir.ConfigDir == "" {
		dir.SetConfigDir()
	}

	exportsFile := filepath.Join(dir.ConfigDir, "zsh", "exports.zsh")
	downloadURL := download.GetDownloadURL("zsh", "exports.zsh")

	err := download.DownloadFiles(downloadURL, exportsFile)
	if err != nil {
		message.ErrorMessage("No se pudo descargar el archivo exports.zsh correctamente", err)
	}
}

func getZshrc() {
	if dir.UserDir == "" {
		dir.SetUserDir()
	}

	zshrcFile := filepath.Join(dir.UserDir, ".zshrc")
	downloadURL := download.GetDownloadURL(".zshrc")

	err := download.DownloadFiles(downloadURL, zshrcFile)
	if err != nil {
		message.ErrorMessage("No se pudo descargar el archivo .zshrc correctamente", err)
	}
}

func getAlias() {
	if dir.ConfigDir == "" {
		dir.SetConfigDir()
	}

	zshDir := filepath.Join(dir.ConfigDir, "zsh")
	aliasFile := filepath.Join(zshDir, "alias.zsh")
	downloadURL := download.GetDownloadURL("zsh", "alias.zsh")

	err := os.MkdirAll(zshDir, 0o755)
	if err != nil {
		message.ErrorMessage("Error al crear la carpeta para guardar el alias", err)
	}

	err = download.DownloadFiles(downloadURL, aliasFile)
	if err != nil {
		message.ErrorMessage("No se pudo descargar el archivo alias.zsh correctamente", err)
	}
}
