package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// ClearCommand clean command struct
type ClearCommand struct {
	BaseCommand
}

// Init ClearCommand
func (cl *ClearCommand) Init() {
	cl.command = &cobra.Command{
		Use:   "clear",
		Short: "Clear lapsed context, cluster and user",
		Long:  "Clear lapsed context, cluster and user",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("赶赶单单啦！")
			return nil
		},
	}

}
