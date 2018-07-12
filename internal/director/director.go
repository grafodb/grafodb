package director

// import (
// 	"net"
// 	"path/filepath"
// 	"sync"
//
// 	"github.com/grafodb/grafodb/internal/cluster"
// 	"github.com/rs/zerolog"
// )
//
// type Config struct {
// 	DataDir           string
// 	RaftAdvertiseAddr *net.TCPAddr
// 	RaftBindAddr      *net.TCPAddr
// 	RaftDebug         bool
// }
//
// func NewDirector(logger *zerolog.Logger, cluster cluster.Cluster, config *Config) Director {
// 	l := logger.With().Str("component", "director").Logger()
// 	return &director{
// 		cluster:           cluster,
// 		dataDir:           filepath.Join(config.DataDir, "director"),
// 		logger:            &l,
// 		raftAdvertiseAddr: config.RaftAdvertiseAddr,
// 		raftBindAddr:      config.RaftBindAddr,
// 		raftDebug:         config.RaftDebug,
// 	}
// }
//
// type Director interface {
// 	Serve(stopWg *sync.WaitGroup, shutdownCh <-chan struct{}, errCh chan<- error)
// }
//
// type director struct {
// 	cluster           cluster.Cluster
// 	dataDir           string
// 	logger            *zerolog.Logger
// 	raftAdvertiseAddr *net.TCPAddr
// 	raftBindAddr      *net.TCPAddr
// 	raftDebug         bool
// }
//
// func (d *director) Serve(stopWg *sync.WaitGroup, shutdownCh <-chan struct{}, errCh chan<- error) {
// 	stopWg.Add(1)
// 	go d.startRaft(stopWg, shutdownCh, errCh)
// }
//
// func (d *director) startRaft(stopWg *sync.WaitGroup, shutdownCh <-chan struct{}, errCh chan<- error) {
// 	defer stopWg.Done()
// }
