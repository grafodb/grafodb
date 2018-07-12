package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print version information",
		RunE:  runVersion,
	}
	rootCmd.AddCommand(versionCmd)
}

func runVersion(_ *cobra.Command, _ []string) error {
	fmt.Printf("GrafoDB - v%s [build=%s]\n", appVersion, appBuild)
	return nil
}
