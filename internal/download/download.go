// Package download it's help you download configuration file
package download

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

const ConfigURL = "https://raw.githubusercontent.com/sikzyo/4k1/main/config"

func DownloadFiles(url, filePath string) error {
	err := os.MkdirAll(filepath.Dir(filePath), 0o755)
	if err != nil {
		return err
	}

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("la descarga falló, el servidor respondió: %s", resp.Status)
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer func() {
		_ = file.Close()
	}()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
