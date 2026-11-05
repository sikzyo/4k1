package config

import (
	"fmt"
	"net/url"
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
	// FIXME: encontrar una mejor manera de generar las URLs dinámicas
	baseURL, _ := url.Parse(download.ConfigURL)
	downloadURL := baseURL.JoinPath("eza", "theme.yaml")

	err := download.DownloadFiles(downloadURL.String(), ezaThemeFile)
	if err != nil {
		message.ErrorMessage("No se pudo descargar el tema para eza", err)
	}
}
