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
	Short: "Prints a list of templates installed in the $STENCIL/templates directory.",
	Long:  `Prints a list of templates installed in the $STENCIL/templates directory.`,
	RunE:  list,
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
