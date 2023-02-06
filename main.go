package main

import (
	"fmt"
	"os"

	"github.com/lantian316/stonetools/cmd"
)

func main() {
	baseCommand := cmd.NewBaseCommand()
	if err := baseCommand.CobraCmd().Execute(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}
