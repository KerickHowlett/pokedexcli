package repl

import (
	"bufio"
	"os"
)

// NewREPL creates a new instance of the REPL struct with the provided options.
//
// It accepts variadic REPLOption as parameters.
//
// The returned REPL object is used to start the Read-Eval-Print Loop (REPL).
//
// Parameters:
//   - options: The options to configure the REPL.
//
// Returns:
//   - A new instance of the REPL struct.
//
// Example usage:
//
//	repl := NewREPL(WithCommandExecutor(commandExecutor))
//
// Now the repl can be used to start the REPL.
func NewREPL(options ...REPLOption) *REPL {
	repl := &REPL{
		prompt:  defaultPrompt,
		scanner: bufio.NewScanner(os.Stdin),
	}

	for _, option := range options {
		option(repl)
	}

	if repl.execute == nil {
		panic("[ERROR] execute is required. Use WithCommandExecutor to set it in REPL.")
	}

	return repl
}