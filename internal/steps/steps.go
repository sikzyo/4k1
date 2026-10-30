// Package steps is to perform the installation steps in order
package steps

import "github.com/sikzyo/4k1/internal/config"

type step struct {
	name    string
	execute func()
}

var gitEmail string

var AllSteps = []step{
	{name: "Xcode", execute: Xcode},
	{name: "Homebrew", execute: homebrew},
	{name: "Git", execute: git},
	{name: "SSH", execute: ssh},
	{name: "Fonts", execute: fonts},
	{name: "Node", execute: node},
	{name: "Go", execute: golang},
	{name: "Terminal apps", execute: cli},
	{name: "Apply Configuration", execute: config.ApplyConfig},
}
