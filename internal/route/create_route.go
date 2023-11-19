package route

import (
	"github.com/william-tome/platform-kit-cli/cmd/config"
	"github.com/william-tome/platform-kit-cli/tools"
)

func CreateRoute(name, serviceName string) error {
	routerDir := "src/router/"
	routesDir := routerDir + "routes/"
	subDirName := tools.HandleStringSuffix(name, "-routes")
	routeSubDir := routesDir + "/" + subDirName
	routePath := routeSubDir + "/index.js"
	indexFilePath := routerDir + "index.js"

	data := config.RouteProps{
		RoutesDir:     routesDir,
		RouteSubDir:   routeSubDir,
		RouteName:     name,
		ServiceName:   serviceName,
		RoutePath:     routePath,
		IndexFilePath: indexFilePath,
	}

	if err := createRoutesDir(&data); err != nil {
		return err
	}

	if err := createRouteFile(&data); err != nil {
		return err
	}

	if err := appendToRouteIndex(&data); err != nil {
		return err
	}

	return nil
}
