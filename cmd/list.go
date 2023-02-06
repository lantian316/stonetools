package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// ListCommand clean command struct
type LIstCommand struct {
	BaseCommand
}

// Init ListCommand
func (ll *ListCommand) Init() {
	ll.command = &cobra.Command{
		Use:   "clear",
		Short: "Clear lapsed context, cluster and user",
		Long:  "Clear lapsed context, cluster and user",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("赶赶单单啦！")
			return nil
		},
	}

}
