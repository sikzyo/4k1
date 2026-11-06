package config

import (
	"fmt"
	"path/filepath"

	"github.com/sikzyo/4k1/internal/dir"
	"github.com/sikzyo/4k1/internal/download"
	"github.com/sikzyo/4k1/internal/message"
)

var (
	fzfDir  string
	fzfFile string
)

func fzfConfig() {
	fmt.Println("✦ fzf ✦")

	setRosepineFzf()
}

func setRosepineFzf() {
	if dir.ConfigDir == "" {
		dir.SetConfigDir()
	}

	fzfDir = filepath.Join(dir.ConfigDir, "fzf")
	fzfFile = filepath.Join(fzfDir, "rose-pine.zsh")
	downloadURL := download.GetDownloadURL("fzf", "rose-pine.zsh")

	err := download.DownloadFiles(downloadURL, fzfFile)
	if err != nil {
		message.ErrorMessage("No se pudo descargar el tema de color para fzf", err)
	}
}
