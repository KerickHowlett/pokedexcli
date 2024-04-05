package repl

import (
	"fmt"
	"strings"
	"testing"

	s "testtools/mocks/scanner"
	"testtools/utils"
)

func TestStartREPL(t *testing.T) {
	setup := func(responseType string) (executedCommand *string, userInput string, output string) {
		executedCommand = new(string)

		scanner := s.NewMockScanner()
		mockedCommandExecutor := func(args string) error {
			*executedCommand = args
			defer scanner.SetIsEnabled(false)

			if responseType == "success" {
				return nil
			}

			if responseType == "error" {
				return fmt.Errorf("error")
			}

			panic("Invalid response type")
		}

		repl := NewREPL(
			WithCommandExecutor(mockedCommandExecutor),
			WithScanner(scanner),
		)

		stdout := utils.NewPrintStorage()
		output = stdout.Capture(func() { repl.RunREPL() })

		userInput = scanner.Text()

		return executedCommand, userInput, output
	}

	t.Run("should execute command and prompt for input until user exits", func(t *testing.T) {
		if executedCommand, expectedCommand, _ := setup("success"); *executedCommand != expectedCommand {
			t.Errorf("Executed command does not match expected command. Got: %s, Expected: %s", *executedCommand, expectedCommand)
		}
	})

	t.Run("should print error message if command execution fails", func(t *testing.T) {
		expectedOutput := "There was an issue with running the mock command. Please try again."
		if _, _, output := setup("error"); !strings.Contains(output, expectedOutput) {
			t.Errorf("Output does not contain expected '%s', but instead got '%s'", expectedOutput, output)
		}
	})
}
