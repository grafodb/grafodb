package graph

import (
	"fmt"
	"path/filepath"

	"github.com/grafodb/grafodb/internal/graph/store"
	"github.com/grafodb/grafodb/internal/util/fs"
	"github.com/rs/zerolog"
)

func NewGraph(logger *zerolog.Logger, dataDir string) (Graph, error) {
	graphDataDir := filepath.Join(dataDir, "graph")
	if err := fs.EnsureDir(graphDataDir); err != nil {
		return nil, fmt.Errorf("error creating graph data dir: %s", graphDataDir)
	}

	graphStore, err := store.NewBadgerStore(logger, graphDataDir, "graph")
	if err != nil {
		return nil, err
	}

	log := logger.With().Str("component", "graph").Logger()
	return &graph{
		logger: &log,
		store:  graphStore,
	}, nil
}

type Graph interface {
}

type graph struct {
	logger *zerolog.Logger
	store  store.GraphStore
}

// func (g *graphServer) Serve(stopWg *sync.WaitGroup, shutdownCh <-chan struct{}, errCh chan<- error) {
// 	stopWg.Add(2)
// 	go g.startServerRPC(stopWg, shutdownCh, errCh)
// 	go g.startServerHTTP(stopWg, shutdownCh, errCh)
// }
//
// func (g *graphServer) startServerHTTP(stopWg *sync.WaitGroup, shutdownCh <-chan struct{}, errCh chan<- error) {
// 	defer stopWg.Done()
//
// 	httpServerAddr := g.httpBindAddr.String()
// 	httpServer := http.Server{
// 		Addr:    httpServerAddr,
// 		Handler: transport.NewTransportHTTP(g.store),
// 	}
//
// 	startErrCh := make(chan error, 1)
// 	doneCh := make(chan struct{}, 1)
//
// 	startFunc := func() {
// 		g.logger.Debug().Msgf("Starting graph HTTP server on %v", httpServerAddr)
// 		listener, err := net.Listen("tcp", httpServerAddr)
// 		if err != nil {
// 			startErrCh <- err
// 			return
// 		}
// 		g.logger.Info().Msgf("Started graph HTTP server on %v", httpServerAddr)
// 		startErrCh <- httpServer.Serve(listener)
// 	}
//
// 	stopFunc := func() {
// 		g.logger.Debug().Msg("Stopping graph HTTP server...")
// 		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 		defer cancel()
//
// 		httpServer.Shutdown(ctx)
// 		close(doneCh)
// 	}
//
// 	go startFunc()
//
// 	select {
// 	case err := <-startErrCh:
// 		errCh <- err
// 	case <-shutdownCh:
// 		stopFunc()
// 	}
//
// 	<-doneCh
// 	g.logger.Info().Msg("Stopped graph HTTP server")
// }
//
// func (g *graphServer) startServerRPC(stopWg *sync.WaitGroup, shutdownCh <-chan struct{}, errCh chan<- error) {
// 	defer stopWg.Done()
//
// 	grpcServerAddr := g.rpcBindAddr.String()
// 	grpcServer := grpc.NewServer()
//
// 	rpc.RegisterGrafoServer(grpcServer, transport.NewTransportRPC(g.store))
//
// 	startErrCh := make(chan error, 1)
// 	doneCh := make(chan struct{}, 1)
//
// 	startFunc := func() {
// 		g.logger.Debug().Msgf("Starting graph RPC server on %v", grpcServerAddr)
// 		listener, err := net.Listen("tcp", grpcServerAddr)
// 		if err != nil {
// 			startErrCh <- err
// 			return
// 		}
// 		g.logger.Info().Msgf("Started graph RPC server on %v", grpcServerAddr)
// 		startErrCh <- grpcServer.Serve(listener)
// 	}
//
// 	stopFunc := func() {
// 		g.logger.Debug().Msg("Stopping graph RPC server...")
// 		grpcServer.GracefulStop()
// 		close(doneCh)
// 	}
//
// 	go startFunc()
//
// 	select {
// 	case err := <-startErrCh:
// 		errCh <- err
// 		close(doneCh)
// 	case <-shutdownCh:
// 		stopFunc()
// 	}
//
// 	<-doneCh
// 	g.logger.Info().Msg("Stopped graph RPC server")
// }
