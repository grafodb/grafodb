package server

import (
	"context"

	"github.com/grafodb/grafodb/internal/graph/store"
	"github.com/grafodb/grafodb/internal/rpc"
)

func NewTransportRPC(store store.GraphStore) rpc.GrafoServer {
	return &rpcServer{
		store: store,
	}
}

type rpcServer struct {
	store store.GraphStore
}

func (s *rpcServer) Send(ctx context.Context, msg *rpc.Message) (*rpc.Message, error) {
	return nil, nil
}
