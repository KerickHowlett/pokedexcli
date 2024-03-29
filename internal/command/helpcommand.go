package command

type HelpCommand struct{}

// Execute is a method of the HelpCommand struct that is responsible for
// executing the actual 'help' command.
//
// It should never be called and will always panic with an error message.
func (c *HelpCommand) Execute() error {
	panic("[ERROR] 'HelpCommand.Execute()' should never be called.")
}

func (c *HelpCommand) GetDescription() string {
	return "Show help for the Pokemon CLI commands."
}

func (c *HelpCommand) GetName() string {
	return "help"
}

func (c *HelpCommand) PrintHelp() {
	printSingleCommandHelp(c)
}

// NewHelpCommand creates a new instance of the HelpCommand struct.
//
// Returns:
//   - A new instance of the HelpCommand struct.
//
// Example usage:
//
//	command := NewHelpCommand()
func NewHelpCommand() *HelpCommand {
	return &HelpCommand{}
}
