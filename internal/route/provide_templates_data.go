package route

import (
	"strings"

	templates "github.com/william-tome/platform-kit-cli/assets/templates/route"
	"github.com/william-tome/platform-kit-cli/cmd/config"
	"github.com/william-tome/platform-kit-cli/tools"
)

func provideRouteTemplateData(d *config.RouteProps) *templates.RouteTemplateFields {
	routeSubDirName := tools.HandleStringSuffix(d.RouteName, "-routes")
	servicesName := tools.HandleStringSuffix(d.ServiceName, "-services")
	routeFileName := tools.ConvertKebabToPascal(d.RouteName)
	routeServices := tools.ConvertKebabToPascal(servicesName)
	routeServicesPath := "@/services/" + servicesName
	routeMethodName := tools.ConvertKebabToCamel(routeSubDirName)
	componentSubDir := "@views/" + routeFileName
	componentPath := componentSubDir + "/" + routeFileName + "View.vue"

	servicesToService := strings.ReplaceAll(servicesName, "services", "service")
	serviceMenthodName := tools.ConvertKebabToCamel(servicesToService)

	data := templates.RouteTemplateFields{
		RouteName:         d.RouteName,
		RouteServices:     routeServices,
		RouteServicesPath: routeServicesPath,
		RouteMethodName:   routeMethodName,
		ComponentPath:     componentPath,
		ServiceMethodName: serviceMenthodName,
	}

	return &data
}
