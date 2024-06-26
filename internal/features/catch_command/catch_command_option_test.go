package catch_command

import (
	bpc "bills_pc"
	p "pokemon"
	qec "query_fetch/query_cache/cache_eviction_config"
	f "test_tools/fixtures"
	"test_tools/utils"
	"testing"
)

func TestWithApiEndpoint(t *testing.T) {
	runWithAPIEndpoint := func() (command *CatchCommand) {
		command = &CatchCommand{}
		WithAPIEndpoint(f.APIEndpoint)(command)
		return command
	}

	t.Run("should set the apiEndpoint field", func(t *testing.T) {
		if command := runWithAPIEndpoint(); command.apiEndpoint != f.APIEndpoint {
			t.Errorf("expected apiEndpoint to be %s, got %s", f.APIEndpoint, command.apiEndpoint)
		}
	})
}

func TestWithCatchPokemon(t *testing.T) {
	runWithCatchPokemon := func() (command *CatchCommand, catchPokemon CatchPokemonFunc) {
		command = &CatchCommand{}
		catchPokemon = func(url string, ec ...*qec.QueryEvictionConfig) (query *p.Pokemon, err error) {
			return &p.Pokemon{}, nil
		}
		WithCatchPokemon(catchPokemon)(command)
		return command, catchPokemon
	}

	t.Run("should set the catchPokemon field", func(t *testing.T) {
		command, catchPokemon := runWithCatchPokemon()
		utils.ExpectSameEntity(t, catchPokemon, command.catchPokemon, "catchPokemon")
	})
}

func TestWithCheckYourLuck(t *testing.T) {
	runWithCheckYourLuck := func() (command *CatchCommand, checkYourLuck CheckYourLuckFunc) {
		command = &CatchCommand{}
		checkYourLuck = func(_ int) int { return 100 }
		WithCheckYourLuck(checkYourLuck)(command)
		return command, checkYourLuck
	}

	t.Run("should set the checkYourLuck field", func(t *testing.T) {
		command, checkYourLuck := runWithCheckYourLuck()
		utils.ExpectSameEntity(t, checkYourLuck, command.checkYourLuck, "checkYourLuck")
	})
}

func TestWithDescription(t *testing.T) {
	runWithDescription := func() (command *CatchCommand, description string) {
		command = &CatchCommand{}
		description = "Catch Description"
		WithDescription(description)(command)
		return command, description
	}

	t.Run("should set the description field", func(t *testing.T) {
		if command, description := runWithDescription(); command.description != description {
			t.Errorf("expected description to be %s, got %s", description, command.description)
		}
	})
}

func TestWithDifficulty(t *testing.T) {
	runWithDifficulty := func() (actual int, expected int) {
		expected = 100
		command := &CatchCommand{}
		WithDifficulty(expected)(command)
		actual = command.difficulty
		return actual, expected
	}

	t.Run("should set the difficulty field", func(t *testing.T) {
		if actual, expected := runWithDifficulty(); actual != expected {
			t.Errorf("expected difficulty to be '%d', got %d", expected, actual)
		}
	})
}

func TestWithEscapedNotification(t *testing.T) {
	runWithEscapedNotification := func() (command *CatchCommand, escapedNotification string) {
		command = &CatchCommand{}
		escapedNotification = "escaped"
		WithEscapedNotification(escapedNotification)(command)
		return command, escapedNotification
	}

	t.Run("should set the escapedNotification field", func(t *testing.T) {
		if command, escapedNotification := runWithEscapedNotification(); command.escapedNotification != escapedNotification {
			t.Errorf("expected escapedNotification to be %s, got %s", escapedNotification, command.escapedNotification)
		}
	})
}

func TestWithEvictionConfig(t *testing.T) {
	runWithEvictionConfig := func() (command *CatchCommand, evictionConfig *qec.QueryEvictionConfig) {
		command = &CatchCommand{}
		evictionConfig = &qec.QueryEvictionConfig{}
		WithEvictionConfig(evictionConfig)(command)
		return command, evictionConfig
	}

	t.Run("should set the evictionConfig field", func(t *testing.T) {
		if command, evictionConfig := runWithEvictionConfig(); command.ec != evictionConfig {
			t.Errorf("expected evictionConfig to be %v, got %v", evictionConfig, command.ec)
		}
	})

}

func TestWithName(t *testing.T) {
	runWithName := func() (command *CatchCommand) {
		command = &CatchCommand{}
		WithName(f.PokemonName)(command)
		return command
	}

	t.Run("should set the name field", func(t *testing.T) {
		if command := runWithName(); command.name != f.PokemonName {
			t.Errorf("expected name to be %s, got %s", f.PokemonName, command.name)
		}
	})
}

func TestWithNoEnteredArgsErrorMessage(t *testing.T) {
	runWithNoEnteredArgsErrorMessage := func() (command *CatchCommand, noEnteredArgsErrorMessage string) {
		command = &CatchCommand{}
		noEnteredArgsErrorMessage = "invalid args"
		WithNoEnteredArgsErrorMessage(noEnteredArgsErrorMessage)(command)
		return command, noEnteredArgsErrorMessage
	}

	t.Run("should set the noEnteredArgsErrorMessage field", func(t *testing.T) {
		if command, noEnteredArgsErrorMessage := runWithNoEnteredArgsErrorMessage(); command.noEnteredArgsErrorMessage != noEnteredArgsErrorMessage {
			t.Errorf("expected noEnteredArgsErrorMessage to be %s, got %s", noEnteredArgsErrorMessage, command.noEnteredArgsErrorMessage)
		}
	})
}

func TestWithPC(t *testing.T) {
	runWithPC := func() (command *CatchCommand, pc *bpc.BillsPC) {
		command = &CatchCommand{}
		pc = &bpc.BillsPC{}
		WithPC(pc)(command)
		return command, pc
	}

	t.Run("should set the pc field", func(t *testing.T) {
		if command, pc := runWithPC(); command.pc != pc {
			t.Errorf("expected pc to be %v, got %v", pc, command.pc)
		}
	})
}

func TestWithSuccessfulCatchNotification(t *testing.T) {
	runWithSuccessfulCatchNotification := func() (command *CatchCommand, successfulCatchNotification string) {
		command = &CatchCommand{}
		successfulCatchNotification = "caught"
		WithSuccessfulCatchNotification(successfulCatchNotification)(command)
		return command, successfulCatchNotification
	}

	t.Run("should set the successfulCatchNotification field", func(t *testing.T) {
		if command, successfulCatchNotification := runWithSuccessfulCatchNotification(); command.successfulCatchNotification != successfulCatchNotification {
			t.Errorf("expected successfulCatchNotification to be %s, got %s", successfulCatchNotification, command.successfulCatchNotification)
		}
	})
}

func TestWithThrowBallNotification(t *testing.T) {
	runWithThrowBallNotification := func() (command *CatchCommand, throwBallNotification string) {
		command = &CatchCommand{}
		throwBallNotification = "thrown"
		WithThrowBallNotification(throwBallNotification)(command)
		return command, throwBallNotification
	}

	t.Run("should set the throwBallNotification field", func(t *testing.T) {
		if command, throwBallNotification := runWithThrowBallNotification(); command.throwBallNotification != throwBallNotification {
			t.Errorf("expected throwBallNotification to be %s, got %s", throwBallNotification, command.throwBallNotification)
		}
	})
}
