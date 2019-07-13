package main

import (
	"github.com/Brialius/goenvdir/cmd"
	"log"
)

var (
	version = "dev"
	build   = "local"
)

func init() {
	log.SetFlags(0)
}

func main() {
	log.Printf("goenvdir %s-%s", version, build)
	cmd.Execute()
}
