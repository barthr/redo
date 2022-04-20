package ui

import (
	"errors"
	"fmt"
	"github.com/barthr/redo/repository"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var quitTextStyle = lipgloss.NewStyle().Margin(1, 0, 2, 2)
var infoTextStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("170"))
var warnTextStyle = lipgloss.NewStyle().
	UnsetPadding().
	UnsetMargins().
	Foreground(lipgloss.Color("#ffa500"))

type ConfirmAliasComponent struct {
	textInput textinput.Model
	err       error
	finalized bool
	selected  []*HistoryItem
	quit      bool
	function  string
}

func newConfirmAliasComponent() *ConfirmAliasComponent {
	textInput := textinput.New()
	textInput.Placeholder = ""
	textInput.Focus()
	textInput.CharLimit = 156
	textInput.Width = 20

	return &ConfirmAliasComponent{textInput: textInput, selected: selectionManager.items}
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
			return c, c.createNewFunction
		}
	case errMsg:
		c.err = msg
		c.textInput.Focus()
		return c, nil
	case functionMsg:
		c.finalized = true
		c.function = string(msg)
		return c, tea.Quit
	}

	c.textInput, cmd = c.textInput.Update(msg)
	return c, cmd
}

type functionMsg string

func (c *ConfirmAliasComponent) createNewFunction() tea.Msg {
	aliasName := c.textInput.Value()
	if aliasName == "" || len(c.selected) == 0 {
		return errMsg{err: errors.New("can't add empty alias or empty commands"), aliasName: aliasName}
	}

	exists, err := repository.GetAliasRepository().Exists(aliasName)
	if err == nil && exists {
		return errMsg{err: fmt.Errorf("sorry that aliasName already exists: " + aliasName), aliasName: aliasName}
	}

	var commands []string
	for _, historyItem := range c.selected {
		commands = append(commands, historyItem.Command)
	}

	var function string
	function, err = repository.GetAliasRepository().Create(repository.Alias{
		Name:     aliasName,
		Commands: commands,
	})
	if err != nil {
		return errMsg{err, aliasName}
	}
	return functionMsg(function)
}

func (c *ConfirmAliasComponent) View() string {
	aliasName := c.textInput.Value()
	if c.quit {
		return ""
	}

	var result string
	if c.err != nil {
		result += warnTextStyle.Render(
			fmt.Sprintf("Something failed when trying to create the alias with name %s: %s", c.err.(errMsg).aliasName, c.err.Error()),
		)
	}

	if !c.finalized {
		result += fmt.Sprintf(
			"\n\nWhat’s the name of the alias?\n\n%s\n\n%s",
			c.textInput.View(),
			"(esc to quit)",
		) + "\n"

		return result
	}

	infoText := infoTextStyle.Render(fmt.Sprintf("Successfully added alias with name: %s\nPlease source your alias file to make your alias active in the current shell \n\n$ source $(redo alias-file)", aliasName))

	return quitTextStyle.Render(fmt.Sprintf("%s\n %s", infoText, c.function))
}
