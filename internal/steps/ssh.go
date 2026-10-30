package steps

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/sikzyo/4k1/internal/dir"
	"github.com/sikzyo/4k1/internal/download"
	"github.com/sikzyo/4k1/internal/message"
	"github.com/sikzyo/4k1/internal/shell"
)

func ssh() {
	fmt.Println("✦ SSH ✦")

	err := validateSSH()
	if err == nil {
		return
	}

	generateSSH()
	sshAgent()
}

func validateSSH() error {
	var err error

	if dir.UserDir == "" {
		dir.SetUserDir()
	}

	file := filepath.Join(dir.UserDir, ".ssh", "github")
	_, err = os.Stat(file)
	if err != nil {
		return err
	}

	return nil
}

func generateSSH() {
	sshFile := filepath.Join(dir.UserDir, ".ssh", "github")

	if gitEmail == "" {
		setGitEmail()
	}

	err := shell.RunCommand(false, "ssh-keygen", "-t", "ed25519", "-C", gitEmail, "-f", sshFile)
	if err != nil {
		message.ErrorMessage("Error al genera clave SSH", err)
	}
}

func sshAgent() {
	sshDir := filepath.Join(dir.UserDir, ".ssh")
	configFile := filepath.Join(sshDir, "config")
	privateFile := filepath.Join(sshDir, "github")

	err := os.MkdirAll(sshDir, 0o700)
	if err != nil {
		message.ErrorMessage("Error al crear el directorio .ssh", err)
	}

	err = download.DownloadFiles("https://raw.githubusercontent.com/sikzyo/4k1/main/config/ssh/config", configFile)
	if err != nil {
		message.ErrorMessage("Error al descargar archivo de configuración", err)
	}

	err = shell.RunCommand(false, "ssh-add", "--apple-use-keychain", privateFile)
	if err != nil {
		message.ErrorMessage("Error al agregar la clave SSH", err)
	}
}
