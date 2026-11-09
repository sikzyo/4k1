package config

import (
	"fmt"
	"path/filepath"

	"github.com/sikzyo/4k1/internal/dir"
	"github.com/sikzyo/4k1/internal/download"
	"github.com/sikzyo/4k1/internal/message"
)

func lazygitConfig() {
	fmt.Println("✦ Lazygit ✦")

	setLazygitConfig()
}

func setLazygitConfig() {
	if dir.ConfigDir == "" {
		dir.SetConfigDir()
	}

	lazygitConfigFile := filepath.Join(dir.ConfigDir, "lazygit", "config.yml")
	downloadURL := download.GetDownloadURL("lazygit", "config.yml")

	if err := download.DownloadFiles(downloadURL, lazygitConfigFile); err != nil {
		message.ErrorMessage("No se pudo descargar la configuración de Lazygit", err)
	}
}
