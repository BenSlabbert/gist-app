package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a gist by id",
	Long:  `Deletes a gist with the provided gist id`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("only provide one argument, the gist id you want to delete")
		}

		id := args[0]
		fmt.Printf("deleting gist: %s", id)
		return api.DeleteGist(id)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
