package main

import (
	"github.com/spf13/cobra"
	"github.com/lantian316/stonetools/cmd"
)

func main() {
	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(cmd.CmdAdd)
	rootCmd.AddCommand(cmd.CmdList)
	rootCmd.Execute()
}
