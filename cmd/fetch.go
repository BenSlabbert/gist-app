package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

// fetchCmd represents the fetch command
var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch your gists",
	Long:  `Fetch your gists returns your gists one page at a time`,
	RunE: func(cmd *cobra.Command, args []string) error {
		gists, err := api.FetchPrivateGists()
		if err != nil {
			return err
		}

		w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
		defer func() {
			_ = w.Flush()
		}()

		_, err = fmt.Fprintln(w, "GistId\tFilename\tType\tSize\tLanguage")
		cobra.CheckErr(err)

		for _, gist := range gists {
			for _, file := range gist.Files {
				_, err = fmt.Fprintf(w, "%s\t%s\t%s\t%d\t%s\n", gist.ID, file.Filename, file.Type, file.Size, file.Language)
				cobra.CheckErr(err)
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(fetchCmd)
}
