package cmd

import (
	"net"
	"strconv"
	"strings"
)

func isValidPort(port int) bool {
	return port > 0 && port <= 65536
}

func parseTCPAddr(addr string, defaultPort int) (*net.TCPAddr, error) {
	if !strings.Contains(addr, ":") {
		addr = net.JoinHostPort(addr, strconv.Itoa(defaultPort))
	}

	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return nil, err
	}
	return tcpAddr, nil
}
