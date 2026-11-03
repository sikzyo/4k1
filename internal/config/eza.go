package config

import (
	"fmt"
	"path/filepath"

	"github.com/sikzyo/4k1/internal/dir"
	"github.com/sikzyo/4k1/internal/download"
	"github.com/sikzyo/4k1/internal/message"
)

var (
	ezaDir       string
	ezaThemeFile string
)

func ezaConfig() {
	fmt.Println("✦ eza ✦")

	setThemeEza()
}

func setThemeEza() {
	if dir.ConfigDir == "" {
		dir.SetConfigDir()
	}

	ezaDir = filepath.Join(dir.ConfigDir, "eza")
	ezaThemeFile = filepath.Join(ezaDir, "theme.yaml")

	err := download.DownloadFiles("https://raw.githubusercontent.com/sikzyo/4k1/heads/main/config/eza/theme.yaml", ezaThemeFile)
	if err != nil {
		message.ErrorMessage("No se pudo descargar el tema para eza", err)
	}
}
