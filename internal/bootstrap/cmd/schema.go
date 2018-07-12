package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	var schemaCmd = &cobra.Command{
		Use:   "schema",
		Short: "Manage the graph schema",
	}
	schemaCmd.PersistentFlags().String(flagServerAddr, defaultServerAddr, flagServerAddrDesc)
	rootCmd.AddCommand(schemaCmd)

	var schemaGetCmd = &cobra.Command{
		Use:   "get",
		Short: "Get a graph schema",
		RunE:  runSchemaGet,
	}
	schemaCmd.AddCommand(schemaGetCmd)

	var schemaSetCmd = &cobra.Command{
		Use:   "set",
		Short: "Set a graph schema",
		RunE:  runSchemaSet,
	}
	schemaSetCmd.Flags().StringP(flagFile, flagFileShort, "", flagFileDesc)
	schemaCmd.AddCommand(schemaSetCmd)

	var schemaListCmd = &cobra.Command{
		Use:   "list",
		Short: "List all graph schemas",
		RunE:  runSchemaList,
	}
	schemaCmd.AddCommand(schemaListCmd)
}

func runSchemaGet(_ *cobra.Command, _ []string) error {
	// FIXME: Missing implementation
	return nil
}

func runSchemaSet(_ *cobra.Command, _ []string) error {
	// filename := viper.GetString(flagFile)
	// if filename == "" {
	// 	return fmt.Errorf("required argument: --file")
	// }
	//
	// data, err := ioutil.ReadFile(filename)
	// if err != nil {
	// 	return err
	// }
	//
	// schema, err := parser.ParseSchema(data)
	// if err != nil {
	// 	return err
	// }
	//
	// jsonData, err := json.Marshal(schema)
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("\n\n%s\n\n", string(jsonData))
	return nil
}

func runSchemaList(_ *cobra.Command, _ []string) error {
	// FIXME: Missing implementation
	return nil
}
