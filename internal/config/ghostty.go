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
	ghosttyConfigDir  string
	ghosttyConfigFile string
)

func ghosttyConfig() {
	fmt.Println("✦ Ghostty ✦")

	if dir.ConfigDir == "" {
		dir.SetConfigDir()
	}

	ghosttyConfigDir = filepath.Join(dir.ConfigDir, "ghostty")
	ghosttyConfigFile = filepath.Join(ghosttyConfigDir, "config.ghostty")
	baseURL, _ := url.Parse(download.ConfigURL)
	downloadURL := baseURL.JoinPath("ghostty", "config.ghostty")

	if err := download.DownloadFiles(downloadURL.String(), ghosttyConfigFile); err != nil {
		message.ErrorMessage("No se pudo descargar la configuración de Ghostty", err)
	}
}
