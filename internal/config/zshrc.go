package config

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/sikzyo/4k1/internal/dir"
	"github.com/sikzyo/4k1/internal/download"
	"github.com/sikzyo/4k1/internal/message"
)

func zshrcConfig() {
	fmt.Println("✦ zsh ✦")

	getZshrc()
	getAlias()
}

func getZshrc() {
	if dir.UserDir == "" {
		dir.SetUserDir()
	}

	zshrcFile := filepath.Join(dir.UserDir, ".zshrc")
	downloadURL := path.Join(download.ConfigURL, ".zshrc")

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
	downloadURL := path.Join(download.ConfigURL, "zsh", "alias.zsh")

	err := os.MkdirAll(zshDir, 0o755)
	if err != nil {
		message.ErrorMessage("Error al crear la carpeta para guardar el alias", err)
	}

	err = download.DownloadFiles(downloadURL, aliasFile)
	if err != nil {
		message.ErrorMessage("No se pudo descargar el archivo alias.zsh correctamente", err)
	}
}
