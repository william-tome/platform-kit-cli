package service

import (
	"strings"

	templates "github.com/william.tome/pk-cli/assets/templates/service"
	"github.com/william.tome/pk-cli/cmd/config"
	"github.com/william.tome/pk-cli/tools"
)

func provideServiceTemplateData(d *config.ServiceProps) *templates.ServiceTemplateFields {
	data := templates.ServiceTemplateFields{
		ServiceBaseUrl:     d.MakeUrlMethodName,
		ServiceBaseUrlPath: "./" + d.MakeUrlFileName,
		ServiceName:        d.ServiceMethodName,
	}

	return &data
}

func provideServiceTestTemplateData(d *config.ServiceProps) *templates.ServiceTestTemplateFields {
	servicePath := d.ServiceSubDir + "/" + d.ServiceName

	data := templates.ServiceTestTemplateFields{
		ServiceName:  d.ServiceMethodName,
		ServicePath:  tools.HandleImportPrefix(servicePath, "@"),
		ServiceTitle: d.TestServiceTitle,
	}

	return &data
}

func provideBaseUrlTemplateData(d *config.ServiceProps) *templates.MakeBaseUrlTemplateFields {
	url := strings.ReplaceAll(d.ServiceName, "-", "/")

	data := templates.MakeBaseUrlTemplateFields{
		MethodName: d.MakeUrlMethodName,
		Url:        url,
	}

	return &data
}

func provideIndexTemplateData(d *config.ServiceProps) *templates.IndexTemplateFields {
	data := templates.IndexTemplateFields{
		MethodName: d.ServiceMethodName,
		Url:        "./" + d.ServiceName,
	}

	return &data
}
