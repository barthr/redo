package repository

import (
	"bytes"
	"log"
	"mvdan.cc/sh/v3/syntax"
	"os"
	"strings"
	"text/template"
	"time"
)

var (
	aliasRepository *AliasRepository
)

func GetAliasRepository() *AliasRepository {
	return aliasRepository
}

type Alias struct {
	Name     string
	Commands []string
}

var functionTemplate = template.Must(template.New("function").Parse(`
# Generated on {{ .Timestamp }}.
{{ .Alias.Name }}() { {{ range .Alias.Commands }}
    {{ . }}{{ end }}
}
`))

type AliasRepository struct {
	aliasFile *os.File
}

func Close() {
	if err := aliasRepository.aliasFile.Close(); err != nil {
		log.Fatalf("Failed to close alias file: %v", err)
	}
}

func InitAliasRepository(aliasFile string) {
	file, err := os.OpenFile(aliasFile, os.O_RDWR, 0644)
	if err != nil {
		log.Fatalf("Failed opening alias file: %s: %s", aliasFile, err)
	}
	aliasRepository = &AliasRepository{aliasFile: file}
}

func (ar *AliasRepository) Create(alias Alias) (string, error) {
	function, err := ar.generateFunction(alias)
	if err != nil {
		return "", err
	}
	err = ar.validateFunction(function)
	if err != nil {
		return "", err
	}
	err = functionTemplate.Execute(ar.aliasFile, map[string]interface{}{
		"Alias":     alias,
		"Timestamp": time.Now().Format(time.RFC3339),
	})
	return function, err
}

func (ar *AliasRepository) generateFunction(alias Alias) (string, error) {
	buffer := &bytes.Buffer{}
	err := functionTemplate.Execute(buffer, map[string]interface{}{
		"Alias":     alias,
		"Timestamp": time.Now().Format(time.RFC3339),
	})
	return buffer.String(), err
}

func (ar *AliasRepository) validateFunction(function string) error {
	_, err := syntax.NewParser().Parse(strings.NewReader(function), "")
	return err
}

func (ar *AliasRepository) Exists(aliasName string) (bool, error) {
	declarations, err := ar.functionDeclarations()
	if err != nil {
		return false, err
	}
	for _, declaration := range declarations {
		if declaration == aliasName {
			return true, nil
		}
	}
	return false, nil
}

func (ar *AliasRepository) functionDeclarations() ([]string, error) {
	parser, err := syntax.NewParser().Parse(ar.aliasFile, "")
	if err != nil {
		return nil, err
	}
	var result []string
	syntax.Walk(parser, func(node syntax.Node) bool {
		switch node.(type) {
		case *syntax.FuncDecl:
			decl := node.(*syntax.FuncDecl)
			result = append(result, decl.Name.Value)
		}
		return true
	})
	return result, nil
}
