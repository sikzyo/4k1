// Package dir is for managing user folders
package dir

import (
	"os"
	"path/filepath"

	"github.com/sikzyo/4k1/internal/message"
)

var (
	UserDir   string
	ConfigDir string
)

func SetUserDir() {
	var err error

	UserDir, err = os.UserHomeDir()
	if err != nil {
		message.ErrorMessage("Error al obtener la carpeta del usuario", err)
	}
}

func SetConfigDir() {
	if UserDir == "" {
		SetUserDir()
	}

	ConfigDir = filepath.Join(UserDir, ".config")
}
