package route_test

import (
	"bytes"
	"testing"

	"github.com/william-tome/platform-kit-cli/assets/templates/route"
	"github.com/william-tome/platform-kit-cli/tests/testdata"
)

func TestGenerateRouteTemplate(t *testing.T) {
	expectedTemplate := testdata.ROUTE_TEMPLATE

	data := route.RouteTemplateFields{
		RouteName:         "route_name",
		RouteServices:     "route_services",
		RouteServicesPath: "route_services_path",
		RouteMethodName:   "route_method_name",
		ComponentPath:     "component_path",
		ServiceMethodName: "service_method_name",
	}

	tmpl, err := route.GenerateRouteTemplate()

	if err != nil {
		t.Errorf("route template: %s", err.Error())
	} else {
		var buf bytes.Buffer
		tmpl.Execute(&buf, data)
		result := buf.String()

		if result != expectedTemplate {
			t.Errorf("Expected template: %s, got: %s", expectedTemplate, result)
		}
	}
}
