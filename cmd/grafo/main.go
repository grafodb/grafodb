package main

import (
	"fmt"
	"os"

	"github.com/grafodb/grafodb/internal/bootstrap/cmd"
)

var (
	Build   = "local"
	Version = "X.X.X"
)

func main() {
	if err := cmd.Execute(Build, Version); err != nil {
		fmt.Printf("error: %s\n", err)
		os.Exit(1)
	}
}
