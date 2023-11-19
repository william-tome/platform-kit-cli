package service

import (
	"bytes"
	"testing"

	"github.com/william-tome/platform-kit-cli/assets/templates/service"
	"github.com/william-tome/platform-kit-cli/tests/testdata"
)

func TestGenerateServiceIndexTemplate(t *testing.T) {
	expectedTemplate := testdata.SERVICE_INDEX_TEMPLATE
	expectedError := "expected_error"

	data := service.IndexTemplateFields{
		MethodName: "method_name",
		Url:        "url",
	}

	tmpl, err := service.GenerateIndexTemplate()

	if err != nil {
		if err.Error() != expectedError {
			t.Errorf("Expected error: %s, got: %s", expectedError, err.Error())
		}
	} else {
		var buf bytes.Buffer
		tmpl.Execute(&buf, data)
		result := buf.String()

		if result != expectedTemplate {
			t.Errorf("Expected template: %s, got: %s", expectedTemplate, result)
		}
	}
}
