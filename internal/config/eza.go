package config

import (
	"fmt"
	"path"
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
	downloadURL := path.Join(download.ConfigURL, "eza", "theme.yaml")

	err := download.DownloadFiles(downloadURL, ezaThemeFile)
	if err != nil {
		message.ErrorMessage("No se pudo descargar el tema para eza", err)
	}
}
