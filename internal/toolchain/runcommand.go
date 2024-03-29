package toolchain

import "fmt"

// RunCommand executes the specified command based on the given selection.
//
// It returns an error if the selection is not valid or if there is an error
// executing the command.
//
// Parameters:
//   - selection: The name of the command to execute.
//
// Returns:
//   - An error if the selection is not valid or if there is an error executing
//     the command.
//
// Example usage:
//
//	toolchain := NewToolchain(
//		WithCommand(c.NewExitCommand()),
//		WithCommand(c.NewHelpCommand()),
//	)
//
//	err := toolchain.RunCommand("exit")
func (t *Toolchain) RunCommand(selection string) error {
	if selection == "" {
		return nil
	}

	if selection == "help" {
		t.PrintHelpMessage()
		return nil
	}

	if selectedCommand, commandFound := t.SelectCommand(selection); commandFound {
		return (*selectedCommand).Execute()
	}

	return fmt.Errorf("[ERROR] Command '%s' is not valid", selection)
}
