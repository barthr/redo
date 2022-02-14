package main

import (
	"flag"
	"fmt"
	"github.com/barthr/redo/repository"
	"github.com/barthr/redo/ui"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"log"
	"os"
	"os/exec"
)

var (
	config = new(Config)
)

func main() {
	config.FromEnv()
	if err := config.EnsureAliasFileExists(); err != nil {
		log.Fatalf("Failed to create alias file %s with error: %s", config.AliasPath, err)
	}

	flag.Parse()

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "alias-file":
			fmt.Println(config.AliasPath)
			os.Exit(0)
		case "edit":
			cmd := exec.Command(os.Getenv("EDITOR"), config.AliasPath)
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			_ = cmd.Run()
			os.Exit(0)
		}
	}

	repository.InitHistoryRepository(config.HistoryPath)
	repository.InitAliasRepository(config.AliasPath)
	defer repository.Close()

	history, err := repository.GetHistoryRepository().GetHistory()
	if err != nil {
		log.Fatalf("Failed fetching history: %v", err)
	}

	var historyItems []list.Item
	for _, historyItem := range history {
		historyItems = append(historyItems, ui.NewHistoryItem(historyItem))
	}
	listComponent := ui.NewHistoryItemListComponent(historyItems)

	if err := tea.NewProgram(listComponent).Start(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
