package server

import (
	"github.com/go-chi/chi"
	"github.com/grafodb/grafodb/internal/graph/store"
)

func NewTransportHTTP(store store.GraphStore) chi.Router {
	server := &httpServer{
		store: store,
	}
	return server.Router()
}

type httpServer struct {
	store store.GraphStore
}

func (s *httpServer) Router() chi.Router {
	router := chi.NewRouter()
	// FIXME: bind HTTP
	return router
}
