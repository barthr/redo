package repository

import (
	"bufio"
	"os"
	"strings"
)

var (
	historyRepository *HistoryRepository
)

func GetHistoryRepository() *HistoryRepository {
	return historyRepository
}

type HistoryRepository struct {
	historyPath string
}

func InitHistoryRepository(historyPath string) {
	historyRepository = &HistoryRepository{historyPath: historyPath}
}

func (repository *HistoryRepository) GetHistory() ([]string, error) {
	readFile, err := os.Open(repository.historyPath)
	if err != nil {
		return nil, err
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var result []string
	for fileScanner.Scan() {
		parsedLine := strings.Split(fileScanner.Text(), ";")
		result = append(result, parsedLine[len(parsedLine)-1])
	}
	// reverse array
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return result, nil
}
