package repository

import (
	"github.com/stretchr/testify/assert"
    "os"
	"testing"
)

func TestAliasRepository_CreateInvalidName(t *testing.T) {
	file, err := os.CreateTemp("", "alias_test")
	assert.NoError(t, err)

	InitAliasRepository(file.Name())

	repository := GetAliasRepository()
	_, err = repository.Create(Alias{
		Name:     "i nvalid",
		Commands: []string{"echo test"},
	})

	assert.Error(t, err, "invalid alias name")
}

func TestAliasRepository_CreateValidName(t *testing.T) {
	file, err := os.CreateTemp("", "alias_test")
	assert.NoError(t, err)

	InitAliasRepository(file.Name())

	repository := GetAliasRepository()
	_, err = repository.Create(Alias{
		Name:     "valid",
		Commands: []string{"echo test"},
	})

	assert.NoError(t, err)
}
