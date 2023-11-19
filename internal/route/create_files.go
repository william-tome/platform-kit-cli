package route

import (
	"fmt"
	"os"

	"github.com/william.tome/pk-cli/cmd/config"
)

func createRouteFile(d *config.RouteProps) error {
	_, err := os.Stat(d.RoutePath)

	if os.IsNotExist(err) {
		file, err := os.Create(d.RoutePath)
		if err != nil {
			return fmt.Errorf("create route file: %s", err.Error())
		}

		data := provideRouteTemplateData(d)
		if err := generateRouteTemplate(data, file); err != nil {
			return err
		}

		defer file.Close()
	}

	return nil
}
