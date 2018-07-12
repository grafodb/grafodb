package cmd

import (
	"fmt"

	"github.com/grafodb/grafodb/internal/util/cliutil"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	var queryCmd = &cobra.Command{
		Use:   "query",
		Short: "Run a query on the graph",
		RunE:  runQuery,
	}
	queryCmd.Flags().String(flagServerAddr, defaultServerAddr, flagServerAddrDesc)
	queryCmd.Flags().StringP(flagFile, flagFileShort, "", flagFileDesc)
	rootCmd.AddCommand(queryCmd)
}

func runQuery(_ *cobra.Command, _ []string) error {
	data, err := cliutil.ReadFromFileOrStdin(viper.GetString(flagFile))
	if err != nil {
		return err
	}
	// FIXME: Missing implementation
	fmt.Println(string(data))
	return nil
}
