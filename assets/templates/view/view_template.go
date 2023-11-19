package view

import (
	"fmt"
	"text/template"
)

type ViewTemplateFields struct {
	BlockName   string
	BlockPath   string
	ServiceName string
}

func GenerateViewTemplate() (*template.Template, error) {
	tmpl, err := template.New("view").Parse(`<template>
  <{{ .BlockName }} :{{ .ServiceName }}="{{ .ServiceName }}" />
</template>

<script setup>
	import {{ .BlockName }} from '{{ .BlockPath }}'

  defineProps({
		{{ .ServiceName }}: {
			type: Function,
			required: false
		}
  })
</script>
`)

	if err != nil {
		return nil, fmt.Errorf("view template: %s", err.Error())
	}

	return tmpl, nil
}
