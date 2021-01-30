package cmd

import (
	"os"

	"github.com/holasoyeddy/stencil/internal"
	"github.com/spf13/cobra"
)

// uninstallCmd represents the uninstall command
var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Removes an installed Stencil template",
	Long:  `Removes an installed Stencil template`,
	RunE:  uninstall,
	Args:  cobra.ExactArgs(1),
}

func init() {
	rootCmd.AddCommand(uninstallCmd)
}

func uninstall(cmd *cobra.Command, args []string) error {
	path, err := internal.GetStencilPath()
	if err != nil {
		return err
	}
	err = os.RemoveAll(path + "/templates/" + args[0])
	if err != nil {
		return err
	}
	return nil
}
