package view

import (
	"fmt"
	"text/template"
)

type BlockTemplateFields struct {
	BlockName   string
	ServiceName string
}

func GenerateBlockTemplate() (*template.Template, error) {
	tmpl, err := template.New("block").Parse(`<template>
  <div>
	 Hello from {{ .BlockName }}
  </div>
</template>

<script setup>
	defineOptions({	name: '{{ .BlockName }}'	})
	
  defineProps({
		{{ .ServiceName }}: {
			type: Function,
			required: false
		}
  })
</script>
`)

	if err != nil {
		return nil, fmt.Errorf("block template: %s", err.Error())
	}

	return tmpl, nil
}
