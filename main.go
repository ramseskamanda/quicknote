package main

import (
	"github.com/ramseskamanda/quicknote/cmd"
	"log"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
