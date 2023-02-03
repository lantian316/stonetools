package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var CmdList = &cobra.Command{
	Use:   "add [string to add]",
	Short: "Add a new item",
	Long: `Add is for adding a new item to the list.
For many years people have added items to lists.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Listing: " + args[0])
	},
}
