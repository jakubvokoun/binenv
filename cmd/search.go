package cmd

import (
	"fmt"

	"github.com/devops-works/binenv/internal/app"
	"github.com/spf13/cobra"
)

// localCmd represents the local command
func searchCmd(a *app.App) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "search [term]",
		Short: "Search term in software distributions",
		Long:  `Search a term in distribution names or descriptions.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			verbose, _ := cmd.Flags().GetBool("verbose")
			a.SetVerbose(verbose)

			switch {
			case len(args) > 1:
				return fmt.Errorf("command requires a single argument")
			case len(args) == 0:
				a.Search("")
			default:
				a.Search(args[0])

			}
			return nil
		},
	}

	return cmd
}
