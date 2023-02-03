package cmd

// NewBaseCommand cmd struct
func NewBaseCommand() *BaseCommand {
	cli := NewCli()
	baseCmd := &BaseCommand{
		command: cli.rootCmd,
	}
	baseCmd.AddCommands(
		&VersionCommand{}, // version command
		&ClearCommand{},   // clear command
	)

	return baseCmd
}
