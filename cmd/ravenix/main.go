package main

import (
	"os"

	"github.com/lucas-farias07/ravenix/internal/audit"
)

func main() {
	command := os.Args[1]
	switch command {
	case "audit":
		audit.Run()
	}
}
