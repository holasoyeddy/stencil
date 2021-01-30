package cmd

import (
	"net/url"
	"os"
	"path"

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
	Args:  cobra.ExactArgs(1),
	RunE:  install,
}

func init() {
	rootCmd.AddCommand(installCmd)
	installCmd.Flags().StringP("name", "n", "", "Specifies the name of the new template.")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// installCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// installCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func install(cmd *cobra.Command, args []string) error {

	// get template name
	name, err := cmd.Flags().GetString("name")

	if err != nil {
		return err
	}

	if name == "" {
		url, err := url.Parse(args[0])

		if err != nil {
			return err
		}

		name = path.Base(url.Path)
	}

	// Get STENCIL path
	stencilPath, err := internal.GetStencilPath()

	if err != nil {
		return err
	}

	// Git clone into $STENCIL/templates directory
	_, err = git.PlainClone(stencilPath+"/templates/"+name, false, &git.CloneOptions{
		URL:      args[0],
		Progress: os.Stdout,
	})

	if err != nil {
		return err
	}

	return nil
}
