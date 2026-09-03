package data

import (
	"github.com/sikzyo/4k1/internal/models"
	"github.com/sikzyo/4k1/internal/steps"
)

var AllSteps = []models.StepsModel{
	{Name: "Xcode", Step: steps.InstallXcode},
}
