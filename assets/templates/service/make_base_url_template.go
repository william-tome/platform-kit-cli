package service

import (
	"fmt"
	"text/template"
)

type MakeBaseUrlTemplateFields struct {
	MethodName string
	Url        string
}

func GenerateBaseUrlTemplate() (*template.Template, error) {
	tmpl, err := template.New("baseUrl").Parse(`export const {{ .MethodName }} = () => {
  return '{{ .Url }}'
}
`)

	if err != nil {
		return nil, fmt.Errorf("service base url template: %s", err.Error())
	}

	return tmpl, nil
}
