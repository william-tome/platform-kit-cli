package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/william-tome/platform-kit-cli/internal/route"
)

func createRoute() *cobra.Command {
	var name, serviceName string

	cmd := &cobra.Command{
		Use:     "route",
		Aliases: []string{"r"},
		Short:   "Creates a route in a new directory under the same name",
		Long: `
		- Creates a route in a "route-name-routes" directory under the provided name.
		- Registers the route in index.js file in the same directory.
		- Service flag is required.
		** Composed route names should use kebab-case. **
		`,
		Example: "route --name my-route --service my-service",
		Run: func(cmd *cobra.Command, args []string) {
			err := route.CreateRoute(name, serviceName)
			if err != nil {
				fmt.Printf("create route command: %s", err)
				os.Exit(1)
			} else {
				fmt.Printf("Route '%s' created! \n", name)
			}
		},
	}
	cmd.Flags().StringVarP(&name, "name", "n", "custom", "The name of the route")
	cmd.Flags().StringVarP(&serviceName, "service", "s", "", "The name of the service to be imported")
	cmd.MarkFlagRequired("service")
	return cmd
}
