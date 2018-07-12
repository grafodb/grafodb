package store

import (
	"github.com/rs/zerolog"
)

func NewBadgerStore(logger *zerolog.Logger, dataDir, name string) (GraphStore, error) {
	log := logger.With().
		Str("component", "store").
		Str("engine", "badger").
		Str("name", name).
		Logger()

	s := &badgerStore{
		dataDir: dataDir,
		name:    name,
		logger:  &log,
	}
	if err := s.initialize(); err != nil {
		return nil, err
	}
	return s, nil
}

type badgerStore struct {
	dataDir string
	name    string
	logger  *zerolog.Logger
}

func (s *badgerStore) initialize() error {
	s.logger.Info().Msg("Initializing graph storage")
	return nil
}
