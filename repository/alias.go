package repository

import (
	"bytes"
	_ "embed"
	"fmt"
	"log"
	"mvdan.cc/sh/v3/syntax"
	"os"
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
{{ .Alias.Name }}() {
    {{ range .Alias.Commands }}
    {{ . }}{{ end }}
}
`))

type AliasRepository struct {
	aliasFile *os.File
}

func Close() {
	aliasRepository.aliasFile.Close()
}

func InitAliasRepository(aliasFile string) {
	file, err := os.OpenFile(aliasFile, os.O_RDWR, 0644)
	if err != nil {
		log.Fatalf("Failed opening alias file: %s: %s", aliasFile, err)
	}
	aliasRepository = &AliasRepository{aliasFile: file}
}

func (ar *AliasRepository) Create(alias Alias) error {
	err := ar.validateFunction(alias)
	if err != nil {
		return err
	}
	err = functionTemplate.Execute(ar.aliasFile, map[string]interface{}{
		"Alias":     alias,
		"Timestamp": time.Now().Format(time.RFC3339),
	})
	return err
}

func (ar *AliasRepository) validateFunction(alias Alias) error {
	buffer := &bytes.Buffer{}
	err := functionTemplate.Execute(buffer, map[string]interface{}{
		"Alias":     alias,
		"Timestamp": time.Now().Format(time.RFC3339),
	})
	if err != nil {
		return err
	}
	_, err = syntax.NewParser().Parse(buffer, "")
	if err != nil {
		return err
	}
	return err
}

func (ar *AliasRepository) Exists(aliasName string) (bool, error) {
	declarations, err := ar.functionDeclarations()
	if err != nil {
		return false, err
	}
	fmt.Println(declarations)
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
