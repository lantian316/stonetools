package main

import (
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(cmd.CmdPrint)
	rootCmd.AddCommand(cmd.CmdList)
	rootCmd.Execute()
}
