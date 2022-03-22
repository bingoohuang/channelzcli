package main

import (
	"os"

	"github.com/bingoohuang/channelzcli/cmd"
)

func main() {
	if err := cmd.NewRootCommand(os.Stdin, os.Stdout).Execute(); err != nil {
		os.Exit(1)
	}
}
