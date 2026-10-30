// Package steps is to perform the installation steps in order
package steps

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
}
