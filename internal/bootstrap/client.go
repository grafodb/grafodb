package bootstrap

import (
	"fmt"
	"net"

	"github.com/grafodb/grafodb/grafo"
)

func NewClientBootstrap(serverAddr *net.TCPAddr) ClientBootstrap {
	return &clientBootstrap{
		serverAddr: serverAddr,
	}
}

type ClientBootstrap interface {
	ClusterMemberList() error
	ClusterStatus() error
}

type clientBootstrap struct {
	serverAddr *net.TCPAddr
}

func (c clientBootstrap) ClusterMemberList() error {
	client, err := grafo.NewClient(c.serverAddr)
	if err != nil {
		return err
	}
	defer client.Close()

	members, err := client.Cluster().MembersList()
	if err != nil {
		return err
	}

	fmt.Printf("%-32s %-16s %s\n", "NAME", "IP", "STATUS")
	for _, m := range members {
		fmt.Printf("%-32s %-16s %s\n", m.ID, m.Addr, m.Status)
	}
	return nil
}

func (c clientBootstrap) ClusterStatus() error {
	return nil
}
