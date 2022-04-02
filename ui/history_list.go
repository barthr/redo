package ui

import (
	"fmt"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"io"
	"strings"
)

var (
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	currentItemStyle  = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("#00ff00"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
)

type HistoryItem struct {
	Command  string
	selected bool
}

func NewHistoryItem(historyEntry string) *HistoryItem {
	historyParts := strings.Split(historyEntry, ";")
	return &HistoryItem{Command: historyParts[len(historyParts)-1]}
}

func (h HistoryItem) FilterValue() string {
	return h.Command
}

type HistoryItemDelegate struct{}

func (h HistoryItemDelegate) Height() int                               { return 1 }
func (h HistoryItemDelegate) Spacing() int                              { return 0 }
func (h HistoryItemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }

func (h HistoryItemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	historyItem, ok := listItem.(*HistoryItem)
	if !ok {
		return
	}
	var str string
	if historyItem.selected {
		str = selectedItemStyle.Render(fmt.Sprintf("%d. %s", index+1, historyItem.Command))
	} else {
		str = fmt.Sprintf("%d. %s", index+1, historyItem.Command)
	}

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s string) string {
			return currentItemStyle.Render("> " + s)
		}
	}

	fmt.Fprint(w, fn(str))
}
