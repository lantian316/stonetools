package cmd

// NewBaseCommand cmd struct
func NewBaseCommand() *BaseCommand {
	cli := NewCli()
	baseCmd := &BaseCommand{
		command: cli.rootCmd,
	}
	baseCmd.AddCommands(
		&AddCommand{}, // add command

		&ClearCommand{}, // clear command
		&ListCommand{}, // clear command
		&CheckCertCommand{}, // clear command

	)

	return baseCmd
}
