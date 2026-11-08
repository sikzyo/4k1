package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/sikzyo/4k1/internal/dir"
	"github.com/sikzyo/4k1/internal/download"
	"github.com/sikzyo/4k1/internal/message"
	"github.com/sikzyo/4k1/internal/shell"
)

var nvimDir string

func nvimConfig() {
	fmt.Println("✦ Neovim ✦")

	setLazyVim()
	setExplorerNvim()
	setThemeNvim()
}

func setLazyVim() {
	if dir.ConfigDir == "" {
		dir.SetConfigDir()
	}

	nvimDir = filepath.Join(dir.ConfigDir, "nvim")
	nvimGitDir := filepath.Join(nvimDir, ".git")

	err := os.RemoveAll(nvimDir)
	if err != nil {
		message.ErrorMessage("No se pudo eliminar la configuración previa de Neovim", err)
	}

	err = shell.RunCommand(false, "git", "clone", "https://github.com/LazyVim/starter", nvimDir)
	if err != nil {
		message.ErrorMessage("No se pudo descargar lazyvim", err)
	}

	err = os.RemoveAll(nvimGitDir)
	if err != nil {
		message.ErrorMessage("No se pudo eliminar la carpeta .git", err)
	}
}

func setExplorerNvim() {
	explorerConfigFile := filepath.Join(nvimDir, "lua", "plugins", "explorer.lua")
	downloadURL := download.GetDownloadURL("nvim", "lua", "plugins", "explorer.lua")

	err := download.DownloadFiles(downloadURL, explorerConfigFile)
	if err != nil {
		message.ErrorMessage("No se pudo descargar la configuración del explorador de Neovim", err)
	}
}

func setThemeNvim() {
	themeConfigFile := filepath.Join(nvimDir, "lua", "plugins", "rose-pine.lua")
	downloadURL := download.GetDownloadURL("nvim", "lua", "plugins", "rose-pine.lua")

	err := download.DownloadFiles(downloadURL, themeConfigFile)
	if err != nil {
		message.ErrorMessage("No se pudo descargar la configuración del tema de Neovim", err)
	}
}
