package ui

import "github.com/charmbracelet/bubbles/list"

type ListAliasComponent struct {
	list     list.Model
	quitting bool
}

func NewListAliasComponent() *ListAliasComponent {
	return &ListAliasComponent{}
}
