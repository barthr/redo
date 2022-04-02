package ui

import (
	"fmt"
	"github.com/barthr/redo/repository"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var quitTextStyle = lipgloss.NewStyle().Margin(1, 0, 2, 2)
var infoTextStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("170"))

type ConfirmAliasComponent struct {
	textInput textinput.Model
	err       error
	finalized bool
	selected  []list.Item
	quit      bool
}

func newConfirmAliasComponent(selected []list.Item) tea.Model {
	textInput := textinput.New()
	textInput.Placeholder = ""
	textInput.Focus()
	textInput.CharLimit = 156
	textInput.Width = 20

	return &ConfirmAliasComponent{textInput: textInput, selected: selected}
}

func (c *ConfirmAliasComponent) Init() tea.Cmd {
	return textinput.Blink
}

func (c *ConfirmAliasComponent) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEscape, tea.KeyCtrlC:
			c.quit = true
			return c, tea.Quit
		case tea.KeyEnter:
			c.finalized = true
			return c, tea.Quit
		}
	case errMsg:
		c.err = msg
		return c, nil
	}

	c.textInput, cmd = c.textInput.Update(msg)
	return c, cmd
}

func (c *ConfirmAliasComponent) View() string {
	if c.quit {
		return ""
	}
	if !c.finalized {
		return fmt.Sprintf(
			"What’s the name of the alias?\n\n%s\n\n%s",
			c.textInput.View(),
			"(esc to quit)",
		) + "\n"
	}

	aliasName := c.textInput.Value()
	if aliasName == "" || len(c.selected) == 0 {
		return quitTextStyle.Render("Can't add empty alias or empty commands")
	}

	exists, err := repository.GetAliasRepository().Exists(aliasName)
	if err == nil && exists {
		return quitTextStyle.Render("Sorry that aliasName already exists: " + aliasName)
	}

	var commands []string
	for _, historyItem := range c.selected {
		commands = append(commands, historyItem.(*HistoryItem).Command)
	}

	var function string
	function, c.err = repository.GetAliasRepository().Create(repository.Alias{
		Name:     aliasName,
		Commands: commands,
	})

	infoText := infoTextStyle.Render(fmt.Sprintf("Successfully added aliasName: %s\nPlease source your alias file to make your alias active in the current shell \n\n$ source $(redo alias-file)", aliasName))
	// source the bash function
	return quitTextStyle.Render(fmt.Sprintf("%s\n %s", infoText, function))
}
