package models

type StepsModel struct {
	Name string
	Step func() error
}
