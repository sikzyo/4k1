package config

import (
	"fmt"
	"path"
	"path/filepath"

	"github.com/sikzyo/4k1/internal/dir"
	"github.com/sikzyo/4k1/internal/download"
	"github.com/sikzyo/4k1/internal/message"
	"github.com/sikzyo/4k1/internal/shell"
)

var (
	batDir        string
	batConfigFile string
	batThemesDir  string
	batThemeFile  string
)

func batConfig() {
	fmt.Println("✦ bat ✦")

	setBatConfig()
	setRosepineBat()
}

func setBatConfig() {
	if dir.ConfigDir == "" {
		dir.SetConfigDir()
	}

	batDir = filepath.Join(dir.ConfigDir, "bat")
	batConfigFile = filepath.Join(batDir, "config")
	downloadURL := path.Join(download.ConfigURL, "bat", "config")

	err := download.DownloadFiles(downloadURL, batConfigFile)
	if err != nil {
		message.ErrorMessage("No se pudo descargar el archivo de configuración para bat", err)
	}
}

func setRosepineBat() {
	if dir.ConfigDir == "" {
		dir.SetConfigDir()
	}

	batThemesDir = filepath.Join(batDir, "themes")
	batThemeFile = filepath.Join(batThemesDir, "rose-pine.tmTheme")
	downloadURL := path.Join(download.ConfigURL, "bat", "themes", "rose-pine.tmTheme")

	err := download.DownloadFiles(downloadURL, batThemeFile)
	if err != nil {
		message.ErrorMessage("No se pudo descargar el tema de color para bat", err)
	}

	err = shell.RunCommand(false, "bat", "cache", "--build")
	if err != nil {
		message.ErrorMessage("No se pudo aplicar el tema descargado", err)
	}
}
