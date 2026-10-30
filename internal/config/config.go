// Package config is used to apply program settings
package config

type config struct {
	name    string
	execute func()
}

var allConfig = []config{
	{name: "zshrc", execute: zshrcConfig},
}
