package director

import (
	"sync"

	"github.com/rs/zerolog"
)

func NewDirector(logger *zerolog.Logger) Director {
	log := logger.With().Str("component", "director").Logger()
	return &director{
		logger: &log,
	}
}

type Director interface {
	StartRaft(stopWg *sync.WaitGroup, shutdownCh <-chan struct{}, errCh chan<- error)
}

type director struct {
	logger *zerolog.Logger
}

func (d *director) StartRaft(stopWg *sync.WaitGroup, shutdownCh <-chan struct{}, errCh chan<- error) {
	defer stopWg.Done()
	// FIXME: Missing implementation
}
