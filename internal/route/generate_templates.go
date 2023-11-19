package route

import (
	"fmt"
	"os"

	templates "github.com/william-tome/platform-kit-cli/assets/templates/route"
)

func generateRouteTemplate(data *templates.RouteTemplateFields, file *os.File) error {
	tmpl, err := templates.GenerateRouteTemplate()

	if err != nil {
		return err
	}

	if err := tmpl.Execute(file, data); err != nil {
		return fmt.Errorf("generate route template: %s", err.Error())
	}

	return nil
}
