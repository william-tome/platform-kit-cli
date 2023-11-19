package service

import (
	"fmt"
	"os"

	"github.com/william.tome/pk-cli/cmd/config"
)

func createServiceFile(d *config.ServiceProps) error {
	_, err := os.Stat(d.ServicePath)

	if os.IsNotExist(err) {
		file, err := os.Create(d.ServicePath)
		if err != nil {
			return fmt.Errorf("create service file: %s", err.Error())
		}

		data := provideServiceTemplateData(d)
		if err := generateServiceTemplate(data, file); err != nil {
			return err
		}

		defer file.Close()
	}

	return nil
}

func createServiceTestFile(d *config.ServiceProps) error {
	_, err := os.Stat(d.TestFilePath)

	if os.IsNotExist(err) {
		file, err := os.Create(d.TestFilePath)
		if err != nil {
			return fmt.Errorf("create service test file: %s", err.Error())
		}

		data := provideServiceTestTemplateData(d)
		if err := generateServiceTestTemplate(data, file); err != nil {
			return err
		}

		defer file.Close()
	}

	return nil
}

func createBaseUrlFile(d *config.ServiceProps) error {
	_, err := os.Stat(d.MakeUrlPath)

	if os.IsNotExist(err) {
		file, err := os.Create(d.MakeUrlPath)
		if err != nil {
			return fmt.Errorf("create base url file: %s", err.Error())
		}

		data := provideBaseUrlTemplateData(d)
		if err := generateBaseUrlTemplate(data, file); err != nil {
			return err
		}

		defer file.Close()
	}

	return nil
}

func createIndexFile(d *config.ServiceProps) error {
	_, err := os.Stat(d.IndexFilePath)

	if os.IsNotExist(err) {
		file, err := os.Create(d.IndexFilePath)
		if err != nil {
			return fmt.Errorf("create index file: %s", err.Error())
		}

		data := provideIndexTemplateData(d)
		if err := generateIndexTemplate(data, file); err != nil {
			return err
		}

		defer file.Close()
	}

	return nil
}
