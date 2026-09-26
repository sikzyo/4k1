// Package menu utilities to render the menu
package menu

type Menu struct {
	Title    string
	Logo     bool
	Options  []MenuOptions
	Fallback string
}
