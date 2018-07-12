package cmd

import (
	"encoding/base64"
	"fmt"

	"github.com/grafodb/grafodb/internal/util/random"
	"github.com/spf13/cobra"
)

func init() {
	var keygenCmd = &cobra.Command{
		Use:   "keygen",
		Short: "Generate a new encryption key",
		RunE:  runKeygen,
	}
	rootCmd.AddCommand(keygenCmd)
}

func runKeygen(_ *cobra.Command, _ []string) error {
	key := base64.StdEncoding.EncodeToString(random.Bytes(16))
	fmt.Println(key)
	return nil
}
