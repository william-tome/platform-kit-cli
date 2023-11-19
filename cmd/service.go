package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/william.tome/pk-cli/internal/service"
)

func createService() *cobra.Command {
	var serviceName, serviceSubDir string
	var append bool

	cmd := &cobra.Command{
		Use:     "service",
		Aliases: []string{"s"},
		Short:   "Creates a service in a new directory under the same name",
		Long: `
		- Creates a service in a "service-name-services" directory under the provided name.
		- Creates the base URL file in the same directory.
		- Registers the service in index.js file in the same directory.
		- If append flag is set, the dir is required.
		** Composed service names should use kebab-case. **
		`,
		Example: `
    // Create service
    service --name my-custom

    // Append service
    service -a --name another-service --dir my-custom-services
    `,
		Run: func(cmd *cobra.Command, args []string) {

			if append {
				if err := service.AppendService(serviceName, serviceSubDir); err != nil {
					fmt.Printf("append service command: %s", err)
					os.Exit(1)
				} else {
					fmt.Printf("Service '%s' created! \n", serviceName)
				}
			} else {
				if err := service.CreateService(serviceName); err != nil {
					fmt.Printf("create service command: %s", err)
					os.Exit(1)
				} else {
					fmt.Printf("Service '%s' created! \n", serviceName)
				}
			}
		},
	}
	cmd.Flags().StringVarP(&serviceName, "name", "n", "custom", "The name of the service")
	cmd.Flags().StringVarP(&serviceSubDir, "dir", "d", "", "The name of the service directory. Eg.: --dir custom-services")
	cmd.Flags().BoolVarP(&append, "append", "a", false, "Append a service to an existing directory under the given name. When appending, the --dir flag is required.")

	if append {
		cmd.MarkFlagRequired("dir")
	}
	return cmd
}
