package main

import (
	"flag"
	"fmt"
	"github.com/barthr/redo/config"
	"github.com/barthr/redo/repository"
	"github.com/barthr/redo/ui"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"log"
	"os"
	"os/exec"
)

var (
	cfg = new(config.Config)
)

const (
	helpText = `
|‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾|
|                               REDO                               |
|__________________________________________________________________|

A Command Line tool to manage all of your shell aliases at one place.

Usage:
  redo (opens up the default browser to create aliases)
  redo [command]

Available Commands:
  help       	Prints out this help text. 
  alias-file 	Prints out the path to the alias file.
  edit 	     	Opens the alias file in your editor (default: %s).
`
)

func main() {
	cfg.FromEnv()
	if err := cfg.EnsureAliasFileExists(); err != nil {
		log.Fatalf("Failed to create alias file %s with error: %s", cfg.AliasPath, err)
	}

	repository.InitHistoryRepository(cfg.HistoryPath)
	repository.InitAliasRepository(cfg.AliasPath)
	defer repository.Close()

	flag.Parse()

	cfg.Validate()

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "help":
			fmt.Fprintf(os.Stdout, helpText, cfg.Editor)
			os.Exit(0)
		case "alias-file":
			fmt.Println(cfg.AliasPath)
			os.Exit(0)
		case "edit":
			openEditor()
			os.Exit(0)
		default:
			log.Fatalf("Command: %s not found", os.Args[1])
		}
	} else {
		history, err := repository.GetHistoryRepository().GetHistory()
		if err != nil {
			log.Fatalf("Failed fetching history: %v", err)
		}

		var historyItems []list.Item
		for _, historyItem := range history {
			historyItems = append(historyItems, ui.NewHistoryItem(historyItem))
		}
		listComponent := ui.NewHistoryItemListComponent(historyItems)

		runTeaProgram(listComponent)
	}
}

func openEditor() {
	cmd := exec.Command(cfg.Editor, cfg.AliasPath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	_ = cmd.Run()
}

func runTeaProgram(root tea.Model) {
	if err := tea.NewProgram(root).Start(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
