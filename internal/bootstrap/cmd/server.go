package cmd

import (
	"fmt"
	"net"
	"strings"

	"github.com/grafodb/grafodb/internal/bootstrap"
	"github.com/hashicorp/go-sockaddr/template"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	var serverCmd = &cobra.Command{
		Use:   "server",
		Short: "Start server",
		RunE:  runServer,
	}

	serverCmd.Flags().String(flagAdvertiseAddr, "", flagAdvertiseAddrDesc)
	serverCmd.Flags().String(flagBindAddr, defaultBindAddr, flagBindAddrDesc)
	serverCmd.Flags().Int(flagBootstrapExpect, defaultBootstrapExpect, flagBootstrapExpectDesc)
	serverCmd.Flags().String(flagClusterID, defaultClusterID, flagClusterIDDesc)
	serverCmd.Flags().String(flagDataDir, defaultDataDir, flagDataDirDesc)
	serverCmd.Flags().Bool(flagDebugRaft, false, flagDebugRaftDesc)
	serverCmd.Flags().Bool(flagDebugSerf, false, flagDebugSerfDesc)
	serverCmd.Flags().Bool(flagNoDirector, false, flagNoDirectorDesc)
	// serverCmd.Flags().String(flagEncrypt, "", flagEncryptDesc)
	serverCmd.Flags().Int(flagHTTPPort, defaultHTTPPort, flagHTTPPortDesc)
	serverCmd.Flags().String(flagJoin, "", flagJoinDesc)
	serverCmd.Flags().Bool(flagNoGraph, false, flagNoGraphDesc)
	serverCmd.Flags().Int(flagRaftPort, defaultRaftPort, flagRaftPortDesc)
	serverCmd.Flags().Int(flagRPCPort, defaultRPCPort, flagRPCPortDesc)
	serverCmd.Flags().Int(flagSerfPort, defaultSerfPort, flagSerfPortDesc)

	serverCmd.Flags().MarkHidden(flagDebugRaft)
	serverCmd.Flags().MarkHidden(flagDebugSerf)

	rootCmd.AddCommand(serverCmd)
}

func runServer(_ *cobra.Command, _ []string) error {
	// Advertise address
	advertiseAddr := viper.GetString(flagAdvertiseAddr)
	if advertiseAddr == "" {
		resolvedAdvertiseAddr, err := template.Parse("{{ GetPrivateIP }}")
		if err != nil {
			return fmt.Errorf("failed to resolve advertise IP address: %v", err)
		}
		advertiseAddr = resolvedAdvertiseAddr
	}

	advertiseIPAddr := net.ParseIP(advertiseAddr)
	if advertiseIPAddr == nil {
		return fmt.Errorf("invalid --%s: %s", flagAdvertiseAddr, advertiseAddr)
	}

	// Bind address
	bindAddr := viper.GetString(flagBindAddr)
	bindIPAddr := net.ParseIP(bindAddr)
	if bindIPAddr == nil {
		return fmt.Errorf("invalid --%s: %s", flagBindAddr, bindAddr)
	}

	// Client HTTP port
	httpPort := viper.GetInt(flagHTTPPort)
	if !isValidPort(httpPort) {
		return fmt.Errorf("invalid --%s: %d", flagHTTPPort, httpPort)
	}

	// RPC port
	rpcPort := viper.GetInt(flagRPCPort)
	if !isValidPort(rpcPort) {
		return fmt.Errorf("invalid --%s: %d", flagRPCPort, rpcPort)
	}

	// Serf port
	serfPort := viper.GetInt(flagSerfPort)
	if !isValidPort(serfPort) {
		return fmt.Errorf("invalid --%s: %d", flagSerfPort, serfPort)
	}

	// Raft port
	raftPort := viper.GetInt(flagRaftPort)
	if !isValidPort(raftPort) {
		return fmt.Errorf("invalid --%s: %d", flagRaftPort, raftPort)
	}

	// Data dir
	dataDir := viper.GetString(flagDataDir)
	if strings.Contains(dataDir, "~") {
		resolvedDataDir, err := homedir.Expand(dataDir)
		if err != nil {
			return fmt.Errorf("failed to expand data directory: %s", dataDir)
		}
		dataDir = resolvedDataDir
	}

	// Join cluster
	joinAddrs := make([]*net.TCPAddr, 0)
	joinArg := viper.GetString(flagJoin)
	for _, addr := range strings.Split(joinArg, ",") {
		tcpAddr, err := parseTCPAddr(addr, defaultSerfPort)
		if err != nil {
			return err
		}
		joinAddrs = append(joinAddrs, tcpAddr)
	}

	enableDirector := !viper.GetBool(flagNoDirector)
	enableGraph := !viper.GetBool(flagNoGraph)

	if !enableDirector && !enableGraph {
		return fmt.Errorf("neither director or graph services enabled for this node")
	}

	if !enableDirector {
		raftPort = -1
	}

	cfg := &bootstrap.ServerBootstrapConfig{
		AdvertiseAddr:     advertiseIPAddr,
		BindAddr:          bindIPAddr,
		BootstrapExpect:   viper.GetInt(flagBootstrapExpect),
		ClusterID:         viper.GetString(flagClusterID),
		DataDir:           dataDir,
		EnableDirector:    enableDirector,
		EnableGraph:       enableGraph,
		HTTPPort:          httpPort,
		JoinAddrs:         joinAddrs,
		RaftPort:          raftPort,
		RaftDebug:         viper.GetBool(flagDebugRaft),
		RPCPort:           rpcPort,
		SerfPort:          serfPort,
		SerfEncryptionKey: viper.GetString(flagEncrypt),
		SerfDebug:         viper.GetBool(flagDebugSerf),
	}
	return bootstrap.NewServerBootstrap(rootLogger, cfg).Start()
}
