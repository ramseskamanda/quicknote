package main

import (
	"github.com/ramseskamanda/quicknote/cmd"
	"log"
)

var (
	version = "dev"
	commit  = "n/a"
)

func main() {
	cmd.Version = version
	cmd.Commit = commit

	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
