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
			return cl.runClear(cmd, args)
		},
		Example: clearExample(),
	}
	cl.command.DisableFlagsInUseLine = true
}

func (cl *ClearCommand) runClear(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		print("len")
	} else {
		print("aaa")
	}
	return nil
}

func clearContext(file string) (bool, error) {
	print("clear")
	return nil
}

func clearExample() string {
	return fmt.Sprintf(`
# Clear lapsed context, cluster and user (default is %s)
kubecm clear
# Customised clear lapsed files
kubecm clear config.yaml test.yaml
`)
}
