package steps

import (
	"github.com/sikzyo/4k1/internal/message"
)

func FullInstall() {
	for _, step := range AllSteps {
		errorMessage, err := step.execute()
		if err != nil {
			message.ErrorMessage(errorMessage, err)
		}
	}
}
