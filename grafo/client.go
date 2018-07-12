package grafo

import (
	"net"

	"github.com/grafodb/grafodb/internal/rpc"
	"google.golang.org/grpc"
)

const (
	ApiV1 = 1
)

// const (
// 	defaultPoolSize = 1
// )

// type ClientOption func(*Client)

// func WithPoolSize(size int) ClientOption {
// 	return func(c *Client) {
// 		c.poolSize = size
// 	}
// }
// opts ...ClientOption
// for _, optFn := range opts {
// optFn(c)
// }

func NewClient(addr *net.TCPAddr) (*Client, error) {
	c := &Client{
		addr: addr,
	}
	c.clusterAPI = &clusterAPI{c}

	if err := c.initConn(); err != nil {
		return nil, err
	}
	return c, nil
}

type Client struct {
	addr       *net.TCPAddr
	conn       *grpc.ClientConn
	poolSize   int
	clusterAPI ClusterAPI
}

func (c *Client) initConn() error {
	dialOpts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(
			grpc.CallCustomCodec(rpc.NewMessageCodec()),
		),
	}
	if conn, err := grpc.Dial(c.addr.String(), dialOpts...); err != nil {
		return err
	} else {
		c.conn = conn
	}
	return nil
}

func (c Client) Cluster() ClusterAPI {
	return c.clusterAPI
}

func (c Client) Close() {
	if c.conn != nil {
		c.conn.Close()
	}
}
