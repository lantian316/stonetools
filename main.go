package main

import (
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(print.CmdPrint)
	rootCmd.AddCommand(add.CmdList)
	rootCmd.Execute()
}
