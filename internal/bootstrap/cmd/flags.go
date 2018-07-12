package cmd

import (
	"fmt"

	"github.com/rs/zerolog"
)

var rootLogger *zerolog.Logger

const (
	flagAdvertiseAddr       = "advertise-addr"
	flagAdvertiseAddrDesc   = "address advertised to other nodes in the cluster (default: private IP)"
	flagBindAddr            = "bind-addr"
	flagBindAddrDesc        = "address to bind services to"
	flagBootstrapExpect     = "bootstrap-expect"
	flagBootstrapExpectDesc = "set the minimum amount of nodes required to boostrap the cluster"
	flagClusterID           = "cluster-id"
	flagClusterIDDesc       = "identifier of the cluster"
	flagDataDir             = "data-dir"
	flagDataDirDesc         = "data directory"
	flagDebug               = "debug"
	flagDebugShort          = "D"
	flagDebugDesc           = "enable debug logging"
	flagDebugRaft           = "debug-raft"
	flagDebugRaftDesc       = "enable raft protocol debug logging"
	flagDebugSerf           = "debug-serf"
	flagDebugSerfDesc       = "enable serf protocol debug logging"
	flagEncrypt             = "encrypt"
	flagEncryptDesc         = "set the gossip encryption key"
	flagFile                = "file"
	flagFileShort           = "f"
	flagFileDesc            = "source file"
	flagHTTPPort            = "http-port"
	flagHTTPPortDesc        = "port used to provide HTTP API"
	flagID                  = "id"
	flagIDDesc              = "identifier of the resource"
	flagJoin                = "join"
	flagJoinDesc            = "comma separated list of node addresses used to join an existing cluster"
	flagLogFormat           = "log-format"
	flagLogFormatShort      = "L"
	flagLogFormatDesc       = "set the log format [console, json]"
	flagNoDirector          = "no-director"
	flagNoDirectorDesc      = "disable director service on this node"
	flagNoGraph             = "no-graph"
	flagNoGraphDesc         = "disable graph service on this node"
	flagSerfPort            = "serf-port"
	flagSerfPortDesc        = "port used for internal communication (serf)"
	flagRaftPort            = "raft-port"
	flagRaftPortDesc        = "port used for internal communication (raft)"
	flagRPCPort             = "rpc-port"
	flagRPCPortDesc         = "port used to provide RPC API"
	flagServerAddr          = "addr"
	flagServerAddrDesc      = "server address"
	flagVerbose             = "verbose"
	flagVerboseDesc         = "enable verbose logging"
)

const (
	defaultBindAddr        = "0.0.0.0"
	defaultBootstrapExpect = 1
	defaultClusterID       = "grafodb"
	defaultSerfPort        = 44014
	defaultRaftPort        = 44024
	defaultDataDir         = "~/.grafodb"
	defaultHTTPPort        = 8444
	defaultRPCPort         = 9444
	defaultLogFormat       = logFormatConsole
	logFormatConsole       = "console"
	logFormatJSON          = "json"
)

var defaultServerAddr = fmt.Sprintf("127.0.0.1:%d", defaultRPCPort)
