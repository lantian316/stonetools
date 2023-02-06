package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// AddCommand add command struct
type AddCommand struct {
	BaseCommand
}



// Init AddCommand
func (ac *AddCommand) Init() {
	ac.command = &cobra.Command{
		Use:     "add",
		Short:   "Add KubeConfig to $HOME/.kube/config",
		Long:    "Add KubeConfig to $HOME/.kube/config",
		Aliases: []string{"a"},
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("何yh大帅比")
			return nil
		},

	}
}
