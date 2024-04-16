package toolchain

import c "command"

// Toolchain represents a collection of commands used in the UI toolchain.
//
// Fields:
//   - commands: A pointer to a map of Commands.
type Toolchain struct {
	// commands is a pointer to a map of Commands.
	commands *c.Commands
}
