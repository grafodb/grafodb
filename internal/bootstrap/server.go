package bootstrap

import (
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/grafodb/grafodb/internal/cluster"
	"github.com/grafodb/grafodb/internal/server"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type ServerBootstrapConfig struct {
	AdvertiseAddr     net.IP
	BindAddr          net.IP
	BootstrapExpect   int
	ClusterID         string
	DataDir           string
	EnableDirector    bool
	EnableGraph       bool
	HTTPPort          int
	JoinAddrs         []*net.TCPAddr
	RaftDebug         bool
	RaftPort          int
	RPCPort           int
	SerfDebug         bool
	SerfEncryptionKey string
	SerfPort          int
}

func NewServerBootstrap(logger *zerolog.Logger, cfg *ServerBootstrapConfig) ServerBootstrap {
	return &serverBootstrap{
		config:     cfg,
		logger:     logger,
		shutdownCh: make(chan struct{}, 1),
	}
}

type ServerBootstrap interface {
	Start() error
	Shutdown()
}

type serverBootstrap struct {
	config       *ServerBootstrapConfig
	logger       *zerolog.Logger
	shutdownLock sync.Mutex
	shutdownCh   chan struct{}
	shutdown     bool
}

func (s *serverBootstrap) Start() error {
	var stopWg sync.WaitGroup
	startErrCh := make(chan error, 1)

	cfg := s.config

	httpAdvertiseAddr := &net.TCPAddr{IP: cfg.AdvertiseAddr, Port: cfg.HTTPPort}
	httpBindAddr := &net.TCPAddr{IP: cfg.BindAddr, Port: cfg.HTTPPort}
	rpcAdvertiseAddr := &net.TCPAddr{IP: cfg.AdvertiseAddr, Port: cfg.RPCPort}
	rpcBindAddr := &net.TCPAddr{IP: cfg.BindAddr, Port: cfg.RPCPort}
	serfAdvertiseAddr := &net.TCPAddr{IP: cfg.AdvertiseAddr, Port: cfg.SerfPort}
	serfBindAddr := &net.TCPAddr{IP: cfg.BindAddr, Port: cfg.SerfPort}

	var raftAdvertiseAddr *net.TCPAddr
	var raftBindAddr *net.TCPAddr
	if cfg.EnableDirector {
		raftAdvertiseAddr = &net.TCPAddr{IP: cfg.AdvertiseAddr, Port: cfg.SerfPort}
		raftBindAddr = &net.TCPAddr{IP: cfg.BindAddr, Port: cfg.SerfPort}
	}

	// cluster
	nodeConfig := &cluster.NodeConfig{
		ClusterID:         cfg.ClusterID,
		DataDir:           cfg.DataDir,
		IsDirector:        cfg.EnableDirector,
		IsGraph:           cfg.EnableGraph,
		HTTPAdvertiseAddr: httpAdvertiseAddr,
		RaftAdvertiseAddr: raftAdvertiseAddr,
		RPCAdvertiseAddr:  rpcAdvertiseAddr,
		SerfAdvertiseAddr: serfAdvertiseAddr,
		SerfBindAddr:      serfBindAddr,
		SerfEncryptionKey: cfg.SerfEncryptionKey,
		SerfDebug:         cfg.SerfDebug,
	}

	c, err := cluster.NewCluster(s.logger, nodeConfig)
	if err != nil {
		return err
	}
	c.Serve(&stopWg, s.shutdownCh, startErrCh)

	s.logger.Debug().Msg("Waiting for cluster manager initialization...")
	<-c.StartedCh()
	s.logger.Debug().Msg("Cluster manager initialized. Bootstrapping server...")

	// join other nodes
	if err := c.Join(cfg.JoinAddrs); err != nil {
		return err
	}

	// server
	serverConfig := &server.Config{
		BootstrapExpect:   s.config.BootstrapExpect,
		DataDir:           cfg.DataDir,
		EnableDirector:    cfg.EnableDirector,
		EnableGraph:       cfg.EnableGraph,
		HTTPBindAddr:      httpBindAddr,
		RaftAdvertiseAddr: raftAdvertiseAddr,
		RaftBindAddr:      raftBindAddr,
		RaftDebug:         cfg.RaftDebug,
		RPCBindAddr:       rpcBindAddr,
	}

	gs, err := server.NewServer(s.logger, c, serverConfig)
	if err != nil {
		return err
	}
	gs.Serve(&stopWg, s.shutdownCh, startErrCh)

	// shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	var startErr error
	select {
	case err := <-startErrCh:
		startErr = err
	case <-quit:
		log.Info().Msg("Initializing server shutdown")
	}
	s.Shutdown()

	s.logger.Debug().Msg("Waiting for shutdown to complete...")
	stopWg.Wait()
	s.logger.Info().Msg("Shutdown completed")

	return startErr
}

func (s *serverBootstrap) Shutdown() {
	s.shutdownLock.Lock()
	defer s.shutdownLock.Unlock()

	if s.shutdown {
		return
	}
	s.shutdown = true
	close(s.shutdownCh)
}
