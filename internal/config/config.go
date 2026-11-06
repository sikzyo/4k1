// Package config is used to apply program settings
package config

type config struct {
	name    string
	execute func()
}

var allConfig = []config{
	{name: "zshrc", execute: zshrcConfig},
	{name: "fzf", execute: fzfConfig},
	{name: "bat", execute: batConfig},
	{name: "micro", execute: microConfig},
	{name: "eza", execute: ezaConfig},
	{name: "ghostty", execute: ghosttyConfig},
}
