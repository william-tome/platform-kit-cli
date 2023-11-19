package view

import (
	"fmt"
	"os"

	templates "github.com/william-tome/platform-kit-cli/assets/templates/view"
)

func generateViewTemplate(data *templates.ViewTemplateFields, file *os.File) error {
	tmpl, err := templates.GenerateViewTemplate()

	if err != nil {
		return err
	}

	if err := tmpl.Execute(file, data); err != nil {
		return fmt.Errorf("generate view template: %s", err.Error())
	}

	return nil
}

func generateBlockTemplate(data *templates.BlockTemplateFields, file *os.File) error {
	tmpl, err := templates.GenerateBlockTemplate()

	if err != nil {
		return err
	}

	if err := tmpl.Execute(file, data); err != nil {
		return fmt.Errorf("generate block template: %s", err.Error())
	}

	return nil
}
