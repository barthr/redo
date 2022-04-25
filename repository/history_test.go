package repository

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetHistoryZsh(t *testing.T) {
	InitHistoryRepository("testdata/zsh.txt")

	repository := GetHistoryRepository()
	history, err := repository.GetHistory()

	assert.NoError(t, err)
	assert.Len(t, history, 3)

	assert.Contains(t, history, "htop")
	assert.Contains(t, history, `echo "test"`)
	assert.Contains(t, history, `docker system prune`)
}

func TestGetHistoryBash(t *testing.T) {
	InitHistoryRepository("testdata/bash.txt")

	repository := GetHistoryRepository()
	history, err := repository.GetHistory()

	assert.NoError(t, err)
	assert.Len(t, history, 3)

	assert.Contains(t, history, "htop")
	assert.Contains(t, history, `echo "test"`)
	assert.Contains(t, history, `docker system prune`)
}
