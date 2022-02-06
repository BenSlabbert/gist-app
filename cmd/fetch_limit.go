package cmd

import (
	"fmt"
	"github.com/BenSlabbert/gist-app/pkg/util"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

// fetchLimitCmd represents the fetchLimit command
var fetchLimitCmd = &cobra.Command{
	Use:   "fetch-limit",
	Short: "Get the current Github Api rate limit",
	Long:  `Get the current Github Api rate limit. This does not count towards the quota`,
	RunE: func(cmd *cobra.Command, args []string) error {
		limit, err := api.GetRateLimit()
		if err != nil {
			return err
		}

		w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
		defer func() {
			_ = w.Flush()
		}()

		_, err = fmt.Fprintln(w, "Limit\tRemaining\tUsed\tReset")
		cobra.CheckErr(err)

		_, err = fmt.Fprintf(w, "%d\t%d\t%d\t%s\n", limit.Resources.Core.Limit, limit.Resources.Core.Remaining, limit.Resources.Core.Used, util.UnixTimestampIntToTime(limit.Resources.Core.Reset))
		cobra.CheckErr(err)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(fetchLimitCmd)
}
