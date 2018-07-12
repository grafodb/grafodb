package cluster

import (
	"fmt"
	"io/ioutil"
	stdlog "log"
	"net"
	"os"
	"path/filepath"
	"sync"

	"strconv"

	"github.com/grafodb/grafodb/internal/util/fs"
	"github.com/grafodb/grafodb/internal/util/random"
	"github.com/hashicorp/serf/serf"
	"github.com/rs/zerolog"
)

const (
	tagClusterID = "cluster-id"
	tagDirector  = "director"
	tagGraph     = "graph"
	tagHTTPAddr  = "graph-http-addr"
	tagRaftAddr  = "node-raft-addr"
	tagRPCAddr   = "graph-rpc-addr"
	tagSerfAddr  = "node-serf-addr"
)

type MemberInfo struct {
	ID         string `json:"id"`
	Addr       string `json:"addr"`
	ClusterID  string `json:"clusterId"`
	IsDirector bool   `json:"isDirector"`
	IsGraph    bool   `json:"isGraph"`
	HTTPAddr   string `json:"httpAddr,omitempty"`
	RaftAddr   string `json:"raftAddr,omitempty"`
	RPCAddr    string `json:"rpcAddr,omitempty"`
	SerfAddr   string `json:"serfAddr,omitempty"`
	Status     string `json:"status"`
}

type NodeConfig struct {
	ClusterID         string
	DataDir           string
	IsDirector        bool
	IsGraph           bool
	HTTPAdvertiseAddr *net.TCPAddr
	RaftAdvertiseAddr *net.TCPAddr
	RPCAdvertiseAddr  *net.TCPAddr
	SerfBindAddr      *net.TCPAddr
	SerfAdvertiseAddr *net.TCPAddr
	SerfEncryptionKey string
	SerfDebug         bool
}

func NewCluster(logger *zerolog.Logger, config *NodeConfig) (Cluster, error) {
	log := logger.With().Str("component", "cluster").Logger()

	serfDataDir := filepath.Join(config.DataDir, "cluster")

	if err := fs.EnsureDir(serfDataDir); err != nil {
		return nil, fmt.Errorf("error creating cluster data dir: %s", serfDataDir)
	}

	return &cluster{
		eventCh:     make(chan serf.Event, 256),
		logger:      &log,
		members:     make(map[string]MemberInfo),
		nodeConfig:  config,
		serfDataDir: serfDataDir,
		startedCh:   make(chan struct{}),
	}, nil
}

type Cluster interface {
	Join(addr *net.TCPAddr) error
	Members() []MemberInfo
	Serf() *serf.Serf
	Serve(stopWg *sync.WaitGroup, shutdownCh <-chan struct{}, errCh chan<- error)
	StartedCh() chan struct{}
}

type cluster struct {
	eventCh     chan serf.Event
	logger      *zerolog.Logger
	memberID    string
	members     map[string]MemberInfo
	membersLock sync.RWMutex
	nodeConfig  *NodeConfig
	serf        *serf.Serf
	serfDataDir string
	startedCh   chan struct{}
}

func (c *cluster) Members() []MemberInfo {
	c.membersLock.RLock()
	defer c.membersLock.RUnlock()

	members := make([]MemberInfo, 0)
	for _, member := range c.members {
		members = append(members, member)
	}
	return members
}

func (c *cluster) Join(addr *net.TCPAddr) error {
	if _, err := c.serf.Join([]string{addr.String()}, true); err != nil {
		return err
	}
	return nil
}

func (c *cluster) Serf() *serf.Serf {
	return c.serf
}

func (c *cluster) Serve(stopWg *sync.WaitGroup, shutdownCh <-chan struct{}, errCh chan<- error) {
	stopWg.Add(1)
	go c.startSerf(stopWg, shutdownCh, errCh)
}

func (c *cluster) StartedCh() chan struct{} {
	return c.startedCh
}

func (c *cluster) startSerf(stopWg *sync.WaitGroup, shutdownCh <-chan struct{}, errCh chan<- error) {
	defer stopWg.Done()

	if err := c.ensureMemberID(); err != nil {
		errCh <- fmt.Errorf("error ensuring cluster member id")
		return
	}

	startErrCh := make(chan error)
	doneCh := make(chan struct{}, 1)

	startFunc := func() {
		defer close(c.startedCh)

		c.logger.Info().Msgf("Starting serf server on %v", c.nodeConfig.SerfBindAddr.String())

		serfConf := serf.DefaultConfig()
		serfConf.Init()
		serfConf.Tags[tagClusterID] = c.nodeConfig.ClusterID
		serfConf.Tags[tagDirector] = strconv.FormatBool(c.nodeConfig.IsDirector)
		serfConf.Tags[tagGraph] = strconv.FormatBool(c.nodeConfig.IsGraph)
		serfConf.Tags[tagHTTPAddr] = c.nodeConfig.HTTPAdvertiseAddr.String()
		serfConf.Tags[tagRPCAddr] = c.nodeConfig.RPCAdvertiseAddr.String()
		serfConf.Tags[tagSerfAddr] = c.nodeConfig.SerfAdvertiseAddr.String()

		if c.nodeConfig.IsDirector {
			serfConf.Tags[tagRaftAddr] = c.nodeConfig.RaftAdvertiseAddr.String()
		}

		serfConf.EventCh = c.eventCh
		serfConf.NodeName = c.memberID
		serfConf.RejoinAfterLeave = false
		serfConf.SnapshotPath = filepath.Join(c.serfDataDir, "serf.snapshot")
		serfConf.MemberlistConfig.AdvertiseAddr = c.nodeConfig.SerfAdvertiseAddr.IP.String()
		serfConf.MemberlistConfig.AdvertisePort = c.nodeConfig.SerfAdvertiseAddr.Port
		serfConf.MemberlistConfig.BindAddr = c.nodeConfig.SerfBindAddr.IP.String()
		serfConf.MemberlistConfig.BindPort = c.nodeConfig.SerfBindAddr.Port
		serfConf.MemberlistConfig.LogOutput = ioutil.Discard

		// if c.serfDebug {
		serfConf.Logger = stdlog.New(c.logger, "", 0)
		// } else {
		// 	serfConf.LogOutput = ioutil.Discard
		// }

		serfServer, err := serf.Create(serfConf)
		if err != nil {
			errCh <- err
			return
		}
		c.serf = serfServer
		c.monitorClusterEvents(shutdownCh)
	}

	stopFunc := func() {
		c.logger.Debug().Msg("Stopping serf server...")
		if c.serf != nil {
			// TODO: Shall we leave the cluster on shutdown?
			c.serf.Leave()
			if err := c.serf.Shutdown(); err != nil {
				errCh <- err
			}
		}
		close(doneCh)
	}

	go startFunc()

	select {
	case err := <-startErrCh:
		errCh <- err
		stopFunc()
	case <-shutdownCh:
		stopFunc()
	}

	<-doneCh
	c.logger.Info().Msg("Stopped serf server")
}

func (c *cluster) ensureMemberID() error {
	if c.memberID != "" {
		return nil
	}

	memberIDFilePath := filepath.Join(c.serfDataDir, "node_id")

	if _, err := os.Stat(memberIDFilePath); err == nil {
		data, err := ioutil.ReadFile(memberIDFilePath)
		if err != nil {
			return err
		}
		c.memberID = string(data)
	} else {
		memberID := fmt.Sprintf("node-%s", random.HexString(8))
		if err := ioutil.WriteFile(memberIDFilePath, []byte(memberID), 0644); err != nil {
			return err
		}
		c.memberID = memberID
	}
	return nil
}

func (c *cluster) monitorClusterEvents(shutdownCh <-chan struct{}) {
	go func() {
		for {
			select {
			case e := <-c.eventCh:
				switch e.EventType() {
				case serf.EventMemberJoin:
					c.addMember(e.(serf.MemberEvent))
				case serf.EventMemberLeave, serf.EventMemberFailed:
					c.removeMember(e.(serf.MemberEvent))
				case serf.EventMemberUpdate:
					c.updateMember(e.(serf.MemberEvent))
					//case serf.EventMemberReap:
					//	//p.localMemberEvent(e.(serf.MemberEvent))
					//case serf.EventMemberUpdate, serf.EventUser, serf.EventQuery:
					//	// Ignore
				default:
					c.logger.Warn().Msgf("unhandled serf event: %#v", e)
				}
			case <-shutdownCh:
				return
			}
		}
	}()
}

func (c *cluster) addMember(ev serf.MemberEvent) {
	for _, member := range ev.Members {
		info, err := c.createMemberInfo(member)
		if err != nil {
			c.logger.Warn().Msgf("Could not add to pool, invalid member: %s", err)
			continue
		}

		c.membersLock.Lock()
		if _, exists := c.members[info.ID]; !exists {
			if c.memberID == info.ID {
				c.logger.Info().Msgf("Adding self to pool: %s [%s]", info.ID, info.Addr)
			} else {
				c.logger.Info().Msgf("Adding member to pool: %s [%s]", info.ID, info.Addr)
			}
			c.members[info.ID] = *info
		}
		c.membersLock.Unlock()
	}
}

func (c *cluster) removeMember(ev serf.MemberEvent) {
	for _, member := range ev.Members {
		info, err := c.createMemberInfo(member)
		if err != nil {
			c.logger.Warn().Msgf("Could not remove from pool, invalid member: %s", err)
			continue
		}

		c.membersLock.Lock()
		if _, exists := c.members[info.ID]; exists {
			c.logger.Info().Msgf("Removing member from pool: %s [%s]", info.ID, info.Addr)
			delete(c.members, info.ID)
		}
		c.membersLock.Unlock()
	}
}

func (c *cluster) updateMember(ev serf.MemberEvent) {
	for _, member := range ev.Members {
		info, err := c.createMemberInfo(member)
		if err != nil {
			c.logger.Warn().Msgf("Could not update in pool, invalid member: %s", err)
			continue
		}

		c.membersLock.Lock()
		if _, exists := c.members[info.ID]; exists {
			c.logger.Info().Msgf("Updating member in pool: %s [%s]", info.ID, info.Addr)
			c.members[info.ID] = *info
		}
		c.membersLock.Unlock()
	}
}

func (c *cluster) createMemberInfo(member serf.Member) (*MemberInfo, error) {
	clusterID := member.Tags[tagClusterID]
	if c.nodeConfig.ClusterID != clusterID {
		return nil, fmt.Errorf("invalid cluster id: %s", clusterID)
	}

	isDirector, err := strconv.ParseBool(member.Tags[tagDirector])
	if err != nil {
		return nil, fmt.Errorf("invalid boolean value for director tag")
	}

	isGraph, err := strconv.ParseBool(member.Tags[tagGraph])
	if err != nil {
		return nil, fmt.Errorf("invalid boolean value for graph tag")
	}

	info := &MemberInfo{
		ID:         member.Name,
		Addr:       member.Addr.String(),
		ClusterID:  member.Tags[tagClusterID],
		IsDirector: isDirector,
		IsGraph:    isGraph,
		HTTPAddr:   member.Tags[tagHTTPAddr],
		RaftAddr:   member.Tags[tagRaftAddr],
		RPCAddr:    member.Tags[tagRPCAddr],
		SerfAddr:   member.Tags[tagSerfAddr],
		Status:     member.Status.String(),
	}
	return info, nil
}

/*
// initKeyring will create a keyring file at a given path.
func initKeyring(path, key string) error {
	var keys []string

	if keyBytes, err := base64.StdEncoding.DecodeString(key); err != nil {
		return fmt.Errorf("Invalid key: %s", err)
	} else if err := memberlist.ValidateKey(keyBytes); err != nil {
		return fmt.Errorf("Invalid key: %s", err)
	}

	// Just exit if the file already exists.
	if _, err := os.Stat(path); err == nil {
		return nil
	}

	keys = append(keys, key)
	keyringBytes, err := json.Marshal(keys)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Dir(path), 0700); err != nil {
		return err
	}

	fh, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer fh.Close()

	if _, err := fh.Write(keyringBytes); err != nil {
		os.Remove(path)
		return err
	}

	return nil
}

// loadKeyringFile will load a gossip encryption keyring out of a file. The file
// must be in JSON format and contain a list of encryption key strings.
func loadKeyringFile(c *serf.Config) error {
	if c.KeyringFile == "" {
		return nil
	}

	if _, err := os.Stat(c.KeyringFile); err != nil {
		return err
	}

	// Read in the keyring file data
	keyringData, err := ioutil.ReadFile(c.KeyringFile)
	if err != nil {
		return err
	}

	// Decode keyring JSON
	keys := make([]string, 0)
	if err := json.Unmarshal(keyringData, &keys); err != nil {
		return err
	}

	// Decode base64 values
	keysDecoded := make([][]byte, len(keys))
	for i, key := range keys {
		keyBytes, err := base64.StdEncoding.DecodeString(key)
		if err != nil {
			return err
		}
		keysDecoded[i] = keyBytes
	}

	// Guard against empty keyring
	if len(keysDecoded) == 0 {
		return fmt.Errorf("no keys present in keyring file: %s", c.KeyringFile)
	}

	// Create the keyring
	keyring, err := memberlist.NewKeyring(keysDecoded, keysDecoded[0])
	if err != nil {
		return err
	}

	c.MemberlistConfig.Keyring = keyring

	// Success!
	return nil
}
*/
