package config

import (
	"fmt"
	"path/filepath"

	"github.com/sikzyo/4k1/internal/dir"
	"github.com/sikzyo/4k1/internal/message"
	"github.com/sikzyo/4k1/internal/shell"
)

func nvimConfig() {
	fmt.Println("✦ Neovim ✦")

	if dir.ConfigDir == "" {
		dir.SetConfigDir()
	}

	nvimDir := filepath.Join(dir.ConfigDir, "nvim")
	if err := shell.RunCommand(true, "git", "clone", "https://github.com/LazyVim/starter", nvimDir); err != nil {
		message.ErrorMessage("No se pudo clonar la configuración de LazyVim", err)
	}
}
