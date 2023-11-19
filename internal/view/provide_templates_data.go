package view

import (
	templates "github.com/william-tome/platform-kit-cli/assets/templates/view"
	"github.com/william-tome/platform-kit-cli/cmd/config"
	"github.com/william-tome/platform-kit-cli/tools"
)

func provideViewTemplateData(d *config.ViewProps) *templates.ViewTemplateFields {
	blockName := tools.ConvertKebabToPascal(d.TemplateName)
	serviceName := tools.ConvertKebabToCamel(d.ServiceName)
	blockPath := "@/templates/" + d.TemplateName

	data := templates.ViewTemplateFields{
		BlockName:   blockName,
		ServiceName: serviceName,
		BlockPath:   blockPath,
	}

	return &data
}

func provideBlockTemplateData(d *config.ViewProps) *templates.BlockTemplateFields {
	serviceName := tools.ConvertKebabToCamel(d.ServiceName)

	data := templates.BlockTemplateFields{
		BlockName:   d.TemplateName,
		ServiceName: serviceName,
	}

	return &data
}
