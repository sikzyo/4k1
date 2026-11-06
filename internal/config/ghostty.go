package config

import (
	"fmt"
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
	downloadURL := download.GetDownloadURL("ghostty", "config.ghostty")

	if err := download.DownloadFiles(downloadURL, ghosttyConfigFile); err != nil {
		message.ErrorMessage("No se pudo descargar la configuración de Ghostty", err)
	}
}
