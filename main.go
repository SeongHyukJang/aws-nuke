package main

import (
	"os"

	"github.com/rebuy-de/aws-nuke/cmd"
)

var (
	// will be overwritten on build
	version = "unknown"
)

type NukeParameters struct {
	ConfigPath string

	Profile         string
	AccessKeyID     string
	SecretAccessKey string

	NoDryRun bool
	Force    bool
}

func main() {
	if err := cmd.NewRootCommand().Execute(); err != nil {
		os.Exit(-1)
	}
}