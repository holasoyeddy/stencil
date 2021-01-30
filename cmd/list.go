package cmd

import (
	"fmt"
	"os"

	"github.com/holasoyeddy/stencil/internal"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: list,
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func list(cmd *cobra.Command, args []string) error {
	path, err := internal.GetStencilPath()
	if err != nil {
		return err
	}
	dir, err := os.Open(path + "/templates")
	if err != nil {
		return err
	}
	directories, err := dir.Readdirnames(0)
	if len(directories) < 1 {
		fmt.Println("no templates installed. Install one with 'stencil install'")
	} else {
		for _, item := range directories {
			fmt.Println(item)
		}
	}
	return nil
}
