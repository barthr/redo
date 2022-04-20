package ui

import (
	"github.com/charmbracelet/bubbles/key"
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

type errMsg struct {
	err       error
	aliasName string
}

func (e errMsg) Error() string { return e.err.Error() }

type HistoryItemListComponent struct {
	list     list.Model
	quitting bool
	selected map[int]*HistoryItem
}

func NewHistoryItemListComponent(items []list.Item) *HistoryItemListComponent {
	l := list.New(items, HistoryItemDelegate{}, defaultWidth, listHeight)
	l.Title = "Which commands would you like to combine?"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(true)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle
	l.AdditionalShortHelpKeys = func() []key.Binding {
		return []key.Binding{
			key.NewBinding(key.WithHelp("space", "toggle item")),
			key.NewBinding(key.WithHelp("enter", "confirm selection")),
		}
	}

	return &HistoryItemListComponent{list: l, selected: map[int]*HistoryItem{}}
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

			if item.isSelected() {
				selectionManager.Remove(item)
			} else {
				selectionManager.Add(item)
			}
		case "enter":
			if len(selectionManager.items) != 0 {
				return newConfirmAliasComponent(), nil
			}
		}
	}

	var cmd tea.Cmd
	h.list, cmd = h.list.Update(msg)
	return h, cmd
}

func (h HistoryItemListComponent) View() string {
	return "\n" + h.list.View()
}
