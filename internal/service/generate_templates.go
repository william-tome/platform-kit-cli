package service

import (
	"fmt"
	"os"

	templates "github.com/william.tome/pk-cli/assets/templates/service"
)

func generateServiceTemplate(data *templates.ServiceTemplateFields, file *os.File) error {
	tmpl, err := templates.GenerateServiceTemplate()

	if err != nil {
		return err
	}

	if err := tmpl.Execute(file, data); err != nil {
		return fmt.Errorf("generate service template: %s", err.Error())
	}

	return nil
}

func generateServiceTestTemplate(data *templates.ServiceTestTemplateFields, file *os.File) error {
	tmpl, err := templates.GenerateServiceTestTemplate()

	if err != nil {
		return err
	}

	if err := tmpl.Execute(file, data); err != nil {
		return fmt.Errorf("generate service test template: %s", err.Error())
	}

	return nil
}

func generateBaseUrlTemplate(data *templates.MakeBaseUrlTemplateFields, file *os.File) error {
	tmpl, err := templates.GenerateBaseUrlTemplate()

	if err != nil {
		return err
	}

	if err := tmpl.Execute(file, data); err != nil {
		return fmt.Errorf("generate base url template: %s", err.Error())
	}

	return nil
}

func generateIndexTemplate(data *templates.IndexTemplateFields, file *os.File) error {
	tmpl, err := templates.GenerateIndexTemplate()

	if err != nil {
		return err
	}

	if err := tmpl.Execute(file, data); err != nil {
		return fmt.Errorf("generate index template: %s", err.Error())
	}

	return nil
}
