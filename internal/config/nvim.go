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
	setExtrasNvim()
	setOptionsNvim()
	setCompletionNvim()
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

func setExtrasNvim() {
	extrasConfigFile := filepath.Join(nvimDir, "lazyvim.json")
	downloadURL := download.GetDownloadURL("nvim", "lazyvim.json")

	err := download.DownloadFiles(downloadURL, extrasConfigFile)
	if err != nil {
		message.ErrorMessage("No se pudo descargar la configuración de extras de Neovim", err)
	}
}

func setOptionsNvim() {
	optionsConfigFile := filepath.Join(nvimDir, "lua", "config", "options.lua")
	downloadURL := download.GetDownloadURL("nvim", "lua", "config", "options.lua")

	err := download.DownloadFiles(downloadURL, optionsConfigFile)
	if err != nil {
		message.ErrorMessage("No se pudo descargar la configuración de opciones de Neovim", err)
	}
}

func setCompletionNvim() {
	completionConfigFile := filepath.Join(nvimDir, "lua", "plugins", "completion.lua")
	downloadURL := download.GetDownloadURL("nvim", "lua", "plugins", "completion.lua")

	err := download.DownloadFiles(downloadURL, completionConfigFile)
	if err != nil {
		message.ErrorMessage("No se pudo descargar la configuración del autocompletado de Neovim", err)
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
