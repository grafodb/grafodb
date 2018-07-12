package server

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"path/filepath"
	"sync"
	"time"

	"github.com/grafodb/grafodb/internal/cluster"
	"github.com/grafodb/grafodb/internal/graph"
	"github.com/grafodb/grafodb/internal/util/fs"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
)

type Config struct {
	BootstrapExpect   int
	DataDir           string
	EnableDirector    bool
	EnableGraph       bool
	HTTPBindAddr      *net.TCPAddr
	RaftAdvertiseAddr *net.TCPAddr
	RaftBindAddr      *net.TCPAddr
	RaftDebug         bool
	RPCBindAddr       *net.TCPAddr
}

func NewServer(logger *zerolog.Logger, cluster cluster.Cluster, config *Config) (Server, error) {
	rootDataDir := filepath.Join(config.DataDir, "data")
	if err := fs.EnsureDir(rootDataDir); err != nil {
		return nil, fmt.Errorf("error creating data dir: %s", rootDataDir)
	}

	var graphDB graph.Graph
	if config.EnableGraph {
		g, err := graph.NewGraph(logger, rootDataDir)
		if err != nil {
			return nil, err
		}
		graphDB = g
	}

	var raftDataDir string
	if config.EnableDirector {
		raftDataDir = filepath.Join(rootDataDir, "raft")
		if err := fs.EnsureDir(raftDataDir); err != nil {
			return nil, fmt.Errorf("error creating raft data dir: %s", raftDataDir)
		}

		// TODO: Initialize director service
	}

	log := logger.With().Str("component", "server").Logger()
	return &grafoServer{
		cluster:      cluster,
		graph:        graphDB,
		httpBindAddr: config.HTTPBindAddr,
		logger:       &log,
		rpcBindAddr:  config.RPCBindAddr,
	}, nil
}

type Server interface {
	Serve(stopWg *sync.WaitGroup, shutdownCh <-chan struct{}, errCh chan<- error)
}

type grafoServer struct {
	cluster           cluster.Cluster
	graph             graph.Graph
	httpBindAddr      *net.TCPAddr
	logger            *zerolog.Logger
	raftAdvertiseAddr *net.TCPAddr
	raftBindAddr      *net.TCPAddr
	raftDebug         bool
	rpcBindAddr       *net.TCPAddr
}

func (g *grafoServer) Serve(stopWg *sync.WaitGroup, shutdownCh <-chan struct{}, errCh chan<- error) {
	stopWg.Add(2)
	go g.startServerRPC(stopWg, shutdownCh, errCh)
	go g.startServerHTTP(stopWg, shutdownCh, errCh)
}

func (g *grafoServer) startRaft(stopWg *sync.WaitGroup, shutdownCh <-chan struct{}, errCh chan<- error) {
	defer stopWg.Done()
	// FIXME: Missing implementation
}

func (g *grafoServer) startServerHTTP(stopWg *sync.WaitGroup, shutdownCh <-chan struct{}, errCh chan<- error) {
	defer stopWg.Done()

	httpServerAddr := g.httpBindAddr.String()
	httpServer := http.Server{
		Addr: httpServerAddr,
		// Handler: transport.NewTransportHTTP(g.store),
	}

	startErrCh := make(chan error, 1)
	doneCh := make(chan struct{}, 1)

	startFunc := func() {
		g.logger.Debug().Msgf("Starting HTTP server on %v", httpServerAddr)
		listener, err := net.Listen("tcp", httpServerAddr)
		if err != nil {
			startErrCh <- err
			return
		}
		g.logger.Info().Msgf("Started HTTP server on %v", httpServerAddr)
		startErrCh <- httpServer.Serve(listener)
	}

	stopFunc := func() {
		g.logger.Debug().Msg("Stopping HTTP server...")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		httpServer.Shutdown(ctx)
		close(doneCh)
	}

	go startFunc()

	select {
	case err := <-startErrCh:
		errCh <- err
	case <-shutdownCh:
		stopFunc()
	}

	<-doneCh
	g.logger.Info().Msg("Stopped HTTP server")
}

func (g *grafoServer) startServerRPC(stopWg *sync.WaitGroup, shutdownCh <-chan struct{}, errCh chan<- error) {
	defer stopWg.Done()

	grpcServerAddr := g.rpcBindAddr.String()
	grpcServer := grpc.NewServer()

	// rpc.RegisterGrafoServer(grpcServer, transport.NewTransportRPC(g.store))

	startErrCh := make(chan error, 1)
	doneCh := make(chan struct{}, 1)

	startFunc := func() {
		g.logger.Debug().Msgf("Starting RPC server on %v", grpcServerAddr)
		listener, err := net.Listen("tcp", grpcServerAddr)
		if err != nil {
			startErrCh <- err
			return
		}
		g.logger.Info().Msgf("Started RPC server on %v", grpcServerAddr)
		startErrCh <- grpcServer.Serve(listener)
	}

	stopFunc := func() {
		g.logger.Debug().Msg("Stopping RPC server...")
		grpcServer.GracefulStop()
		close(doneCh)
	}

	go startFunc()

	select {
	case err := <-startErrCh:
		errCh <- err
		close(doneCh)
	case <-shutdownCh:
		stopFunc()
	}

	<-doneCh
	g.logger.Info().Msg("Stopped RPC server")
}
