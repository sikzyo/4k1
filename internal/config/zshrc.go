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
	getAlias()
}

func getZshrc() {
	if dir.UserDir == "" {
		dir.SetUserDir()
	}

	zshrcFile := filepath.Join(dir.UserDir, ".zshrc")

	err := download.DownloadFiles("https://raw.githubusercontent.com/sikzyo/4k1/heads/main/config/.zshrc", zshrcFile)
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

	err := os.MkdirAll(zshDir, 0o755)
	if err != nil {
		message.ErrorMessage("Error al crear la carpeta para guardar el alias", err)
	}

	err = download.DownloadFiles("https://raw.githubusercontent.com/sikzyo/4k1/heads/main/config/zsh/alias.zsh", aliasFile)
	if err != nil {
		message.ErrorMessage("No se pudo descargar el archivo alias.zsh correctamente", err)
	}
}
