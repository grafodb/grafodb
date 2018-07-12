package cmd

func isValidPort(port int) bool {
	return port > 0 && port <= 65536
}
