package handler

import (
	"context"
	"fmt"

	"github.com/grafodb/grafodb/internal/cluster"
	"github.com/grafodb/grafodb/internal/director"
	"github.com/grafodb/grafodb/internal/graph"
	"github.com/grafodb/grafodb/internal/rpc"
)

type HandleFunc func(version byte, payload []byte) (*rpc.Message, error)

func NewHandler(cluster cluster.Cluster, director director.Director, graph graph.Graph) *Handler {
	h := &Handler{
		cluster:  cluster,
		director: director,
		graph:    graph,
		handlers: make(map[uint16]HandleFunc),
	}
	h.init()
	return h
}

type Handler struct {
	cluster  cluster.Cluster
	director director.Director
	graph    graph.Graph
	handlers map[uint16]HandleFunc
}

func (h *Handler) init() {
	h.handlers[rpc.CmdClusterMemberList] = h.clusterMemberList
	h.handlers[rpc.CmdClusterStatus] = h.clusterStatus
}

func (h Handler) Dispatch(ctx context.Context, msg *rpc.Message) (*rpc.Message, error) {

	funcErrCh := make(chan error, 1)
	funcResultCh := make(chan *rpc.Message, 1)

	defer func() {
		close(funcErrCh)
		close(funcResultCh)
	}()

	go func() {
		req := rpc.GetRootAsRequest(msg.Payload, 0)
		command := req.Command()
		version := req.Version()

		fn, exists := h.handlers[command]
		if !exists {
			funcErrCh <- fmt.Errorf("function handler not found for command: %d", command)
			return
		}

		if result, err := fn(version, req.PayloadBytes()); err != nil {
			funcErrCh <- err
		} else {
			funcResultCh <- result
		}
	}()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case err := <-funcErrCh:
		return nil, err
	case result := <-funcResultCh:
		return result, nil
	}
}
