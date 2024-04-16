package help_command

import "testing"

func TestWithDescription(t *testing.T) {
	runWithDescriptionTest := func() (actual string, expected string) {
		command := &HelpCommand{}
		expected = "Help Description."
		WithDescription(expected)(command)
		return command.description, expected

	}

	t.Run("should set the description of the HelpCommand", func(t *testing.T) {
		if actual, expected := runWithDescriptionTest(); actual != expected {
			t.Errorf("Expected description to be '%s', but got '%s'", expected, actual)
		}
	})
}

func TestWithName(t *testing.T) {
	runWithNameTest := func() (actual string, expected string) {
		command := &HelpCommand{}
		expected = "help-name"
		WithName(expected)(command)
		return command.name, expected

	}

	t.Run("should set the name of the HelpCommand", func(t *testing.T) {
		if actual, expected := runWithNameTest(); actual != expected {
			t.Errorf("Expected name to be '%s', but got '%s'", expected, actual)
		}
	})
}
