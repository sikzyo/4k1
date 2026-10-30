package config

import (
	"github.com/sikzyo/4k1/internal/message"
	"github.com/sikzyo/4k1/internal/shell"
)

// TODO: Se debe de tomar todas las configuraciones y ejecutarlas una por una en orden

func ApplyConfig() {
	for _, config := range allConfig {
		err := shell.RunCommand(false, "clear")
		if err != nil {
			message.ErrorMessage("No se puedo limpiar la pantalla correctamente", err)
		}

		config.execute()
	}
}
