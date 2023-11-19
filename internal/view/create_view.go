package view

import (
	"github.com/william-tome/platform-kit-cli/cmd/config"
	"github.com/william-tome/platform-kit-cli/tools"
)

func CreateView(name, serviceName, templateName string) error {
	viewsDir := "src/views/"
	blocksDir := "src/templates/"

	viewCompName := tools.ConvertKebabToPascal(name)
	viewName := tools.HandleStringSuffix(name, "-view")
	viewSubDir := viewsDir + viewCompName
	templateName = tools.HandleStringSuffix(templateName, "-block")
	blockSubDir := blocksDir + templateName
	serviceName = tools.HandleStringSuffix(serviceName, "-service")
	blockPath := blockSubDir + "/index.vue"
	viewPath := viewSubDir + "/" + viewCompName + "View.vue"

	data := config.ViewProps{
		ViewName:     viewName,
		ViewSubDir:   viewSubDir,
		TemplateName: templateName,
		ServiceName:  serviceName,
		BlockSubDir:  blockSubDir,
		BlockPath:    blockPath,
		ViewPath:     viewPath,
	}

	if err := createViewDir(&data); err != nil {
		return err
	}

	if err := createBlockDir(&data); err != nil {
		return err
	}

	if err := createViewFile(&data); err != nil {
		return err
	}

	if err := createBlockFile(&data); err != nil {
		return err
	}

	return nil
}
