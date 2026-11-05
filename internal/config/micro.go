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
	microDir        string
	microConfigFile string
	microThemeDir   string
	microThemeFile  string
)

func microConfig() {
	fmt.Println("✦ Micro ✦")

	setMicroConfig()
	setThemeMicro()
}

func setMicroConfig() {
	if dir.ConfigDir == "" {
		dir.SetConfigDir()
	}

	microDir = filepath.Join(dir.ConfigDir, "micro")
	microConfigFile = filepath.Join(microDir, "settings.json")
	downloadURL := path.Join(download.ConfigURL, "micro", "settings.json")

	err := download.DownloadFiles(downloadURL, microConfigFile)
	if err != nil {
		message.ErrorMessage("No se pudo descargar la configuración de Micro", err)
	}
}

func setThemeMicro() {
	microThemeDir = filepath.Join(microDir, "colorschemes")
	microThemeFile = filepath.Join(microThemeDir, "rose-pine.micro")
	downloadURL := path.Join(download.ConfigURL, "micro", "colorschemes", "rose-pine.micro")

	err := download.DownloadFiles(downloadURL, microThemeFile)
	if err != nil {
		message.ErrorMessage("No se pudo descargar el tema para micro", err)
	}
}
