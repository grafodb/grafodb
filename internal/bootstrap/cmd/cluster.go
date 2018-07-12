package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	var clusterCmd = &cobra.Command{
		Use:   "cluster",
		Short: "Manage the cluster",
	}
	clusterCmd.Flags().String(flagServerAddr, defaultServerAddr, flagServerAddrDesc)
	clusterCmd.AddCommand(clusterConfigCommand(), clusterMemberListCommand(), clusterStatusCommand())
	rootCmd.AddCommand(clusterCmd)
}

func clusterConfigCommand() *cobra.Command {
	var clusterConfigCmd = &cobra.Command{
		Use:   "config",
		Short: "Manage cluster configuration",
	}

	// var clusterConfigSetCmd = &cobra.Command{
	// 	Use:   "add",
	// 	Short: "Add a new configuration resource",
	// 	RunE:  runClusterConfigCreate,
	// }
	// clusterConfigSetCmd.Flags().String(flagID, "", fmt.Sprintf("%s (%s)", flagIDDesc, "defaults to resource computed hash"))
	// clusterConfigSetCmd.Flags().StringP(flagFile, flagFileShort, "", flagFileDesc)
	// clusterConfigCmd.AddCommand(clusterConfigSetCmd)
	//
	// var clusterConfigGetCmd = &cobra.Command{
	// 	Use:   "get",
	// 	Short: "Get a configuration",
	// 	RunE:  runClusterConfigGet,
	// }
	// clusterConfigGetCmd.Flags().String(flagID, "", flagIDDesc)
	// clusterConfigCmd.AddCommand(clusterConfigGetCmd)
	//
	// var clusterConfigDeleteCmd = &cobra.Command{
	// 	Use:   "delete",
	// 	Short: "Delete a configuration resource",
	// 	RunE:  runClusterConfigDelete,
	// }
	// clusterConfigDeleteCmd.Flags().String(flagID, "", flagIDDesc)
	// clusterConfigCmd.AddCommand(clusterConfigDeleteCmd)
	//
	// var clusterConfigApplyCmd = &cobra.Command{
	// 	Use:   "apply",
	// 	Short: "Apply a configuration resource to the cluster",
	// 	RunE:  runClusterConfigApply,
	// }
	// clusterConfigApplyCmd.Flags().String(flagID, "", flagIDDesc)
	// clusterConfigCmd.AddCommand(clusterConfigApplyCmd)
	//
	// var clusterConfigListCmd = &cobra.Command{
	// 	Use:   "list",
	// 	Short: "List all graph clusterConfigs",
	// 	RunE:  runClusterConfigList,
	// }
	// clusterConfigCmd.AddCommand(clusterConfigListCmd)

	return clusterConfigCmd
}

func clusterMemberListCommand() *cobra.Command {
	var clusterMemberListCmd = &cobra.Command{
		Use:   "members",
		Short: "List all cluster members",
		RunE:  runClusterMemberList,
	}
	return clusterMemberListCmd
}

func clusterStatusCommand() *cobra.Command {
	var clusterStatusCmd = &cobra.Command{
		Use:   "status",
		Short: "Get the cluster status",
		RunE:  runClusterStatus,
	}
	return clusterStatusCmd
}

func runClusterMemberList(_ *cobra.Command, _ []string) error {
	// FIXME: Missing implementation
	return nil
}

func runClusterStatus(_ *cobra.Command, _ []string) error {
	// FIXME: Missing implementation
	return nil
}
