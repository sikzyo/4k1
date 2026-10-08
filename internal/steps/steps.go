// Package steps is to perform the installation steps in order
package steps

type step struct {
	name    string
	execute func() (errorMessaje string, err error)
}

var AllSteps = []step{
	{name: "Xcode", execute: Xcode},
}
