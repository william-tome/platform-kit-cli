package view

import (
	"fmt"
	"os"

	"github.com/william.tome/pk-cli/cmd/config"
)

func createViewFile(d *config.ViewProps) error {
	_, err := os.Stat(d.ViewPath)

	if os.IsNotExist(err) {
		file, err := os.Create(d.ViewPath)
		if err != nil {
			return fmt.Errorf("create view file: %s", err.Error())
		}

		data := provideViewTemplateData(d)
		if err := generateViewTemplate(data, file); err != nil {
			return err
		}

		defer file.Close()
	}

	return nil
}

func createBlockFile(d *config.ViewProps) error {
	_, err := os.Stat(d.BlockPath)

	if os.IsNotExist(err) {
		file, err := os.Create(d.BlockPath)
		if err != nil {
			return fmt.Errorf("create block file: %s", err.Error())
		}

		data := provideBlockTemplateData(d)
		if err := generateBlockTemplate(data, file); err != nil {
			return err
		}

		defer file.Close()
	}

	return nil
}
