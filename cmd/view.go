package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/william.tome/pk-cli/internal/view"
)

func createView() *cobra.Command {
	var name, serviceName, templateName string

	cmd := &cobra.Command{
		Use:     "view",
		Aliases: []string{"v"},
		Short:   "Creates a view in a new directory under the same name",
		Long: `
		- Creates a view in a "ViewName" directory under the provided name.
		- Imports the service.
		- Creates the template in the templates directory and imports it.
		** Composed view names should use kebab-case. **
		`,
		Example: "view --name my-view --service my-service",
		Run: func(cmd *cobra.Command, args []string) {
			err := view.CreateView(name, serviceName, templateName)
			if err != nil {
				fmt.Printf("create view command: %s", err)
				os.Exit(1)
			} else {
				fmt.Printf("View '%s' created! \n", name)
			}
		},
	}
	cmd.Flags().StringVarP(&name, "name", "n", "custom", "The name of the view")
	cmd.Flags().StringVarP(&serviceName, "service", "s", "custom", "The name of the service to be imported")
	cmd.Flags().StringVarP(&templateName, "template", "t", "custom", "The name of the template to be imported")
	return cmd
}
