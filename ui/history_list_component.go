package ui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	defaultWidth = 20
	listHeight   = 30
)

var (
	titleStyle = lipgloss.NewStyle().MarginLeft(2)
)

type errMsg struct{ err error }

func (e errMsg) Error() string { return e.err.Error() }

type HistoryItemListComponent struct {
	list     list.Model
	quitting bool
}

func NewHistoryItemListComponent(items []list.Item) *HistoryItemListComponent {
	l := list.New(items, HistoryItemDelegate{}, defaultWidth, listHeight)
	l.Title = "Which commands would you like to combine?"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(true)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	return &HistoryItemListComponent{list: l}
}

func (h HistoryItemListComponent) Init() tea.Cmd { return nil }
func (h HistoryItemListComponent) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h.list.SetWidth(msg.Width)
		return h, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "ctrl+c":
			h.quitting = true
			return h, tea.Quit

		case " ":
			item := h.list.SelectedItem().(*HistoryItem)
			item.selected = !item.selected
		case "enter":
			return newConfirmAliasComponent(h.selectedItems()), nil
		}
	}

	var cmd tea.Cmd
	h.list, cmd = h.list.Update(msg)
	return h, cmd
}

func (h HistoryItemListComponent) View() string {
	return "\n" + h.list.View()
}

func (h HistoryItemListComponent) selectedItems() []list.Item {
	var result []list.Item
	for _, item := range h.list.Items() {
		if item.(*HistoryItem).selected {
			result = append(result, item)
		}
	}
	return result
}
