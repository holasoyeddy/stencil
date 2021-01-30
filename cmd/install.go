package cmd

import (
	"os"

	"github.com/go-git/go-git/v5"
	_ "github.com/go-git/go-git/v5"
	"github.com/holasoyeddy/stencil/internal"
	"github.com/spf13/cobra"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Downloads and installs a Stencil template.",
	Long:  `Downloads a Stencil template from a git repository and installs it in the $STENCIL/templates path.`,
	RunE:  install,
}

func init() {
	rootCmd.AddCommand(installCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// installCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// installCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func install(cmd *cobra.Command, args []string) error {

	path, err := internal.GetStencilPath()

	if err != nil {
		return err
	}

	_, err = git.PlainClone(path+"/templates/go-git", false, &git.CloneOptions{
		URL:      "https://github.com/go-git/go-git",
		Progress: os.Stdout,
	})

	if err != nil {
		return err
	}

	return nil
}
