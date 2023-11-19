package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/william-tome/platform-kit-cli/internal/create"
)

func createAllFiles() *cobra.Command {
	var name string

	cmd := &cobra.Command{
		Use:     "create",
		Aliases: []string{"c"},
		Short:   "Creates all necessary files for a new view",
		Long: `
		- Creates and registers a new route under the provided name.
		- Creates and registers a new service under the provided name.
		- Creates and registers a new view under the provided name.
		- Creates and registers a new block under the provided name.
		- Creates the service unit test template.

		** Composed names should use kebab-case. **
		`,
		Example: "create --name new-module",
		Run: func(cmd *cobra.Command, args []string) {
			err := create.CreateAllFiles(name)
			if err != nil {
				fmt.Printf("create all files command: %s", err)
				os.Exit(1)
			} else {
				fmt.Printf("Files for '%s' created! \n", name)
			}
		},
	}
	cmd.Flags().StringVarP(&name, "name", "n", "custom", "The name given to the created files")
	return cmd
}
