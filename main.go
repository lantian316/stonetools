package main

import (
	"github.com/lantian316/stonetools/cmd"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(cmd.CmdPrint)
	rootCmd.AddCommand(cmd.CmdList)
	rootCmd.Execute()
}
