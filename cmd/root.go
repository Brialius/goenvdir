package cmd

import (
	"github.com/Brialius/goenvdir/internal"
	"log"
	"os"
)

const usage = `
Usage:
goenvdir dir child
	dir
		directory with files named as env variables
	child
		executable program with all parameters
`

// return special exit code as original envdir tool
func failExit(message string) {
	log.Println("Failed to execute goenvdir")
	log.Println(message)
	os.Exit(111)
}

// Execute root command implementation
func Execute() {
	dir, child := parseArgs(os.Args)
	if exitCode, err := internal.EnvDir(dir, child); err != nil {
		failExit(err.Error())
	} else {
		os.Exit(exitCode)
	}
}

func parseArgs(args []string) (string, []string) {
	args = args[1:]
	if len(args) < 2 {
		log.Println("Required at leas two arguments")
		failExit(usage)
	}
	return args[0], args[1:]
}
