package service

import (
	"fmt"
	"text/template"
)

type IndexTemplateFields struct {
	MethodName string
	Url        string
}

func GenerateIndexTemplate() (*template.Template, error) {
	tmpl, err := template.New("index").Parse(`import { {{ .MethodName }} } from '{{ .Url }}'

export {
	{{ .MethodName }}
}
`)

	if err != nil {
		return nil, fmt.Errorf("service index template: %s", err.Error())
	}

	return tmpl, nil
}
