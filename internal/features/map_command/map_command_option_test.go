package map_command

import (
	"testing"

	pd "map_command/pagination_direction"
	ms "map_command/state"
	qec "query_fetch/query_cache/cache_eviction_config"
	"test_tools/utils"
)

func TestWithCommandDescription(t *testing.T) {
	runWithCommandDescriptionTest := func() (commandDescription string, cmd *MapCommand) {
		cmd = &MapCommand{}
		commandDescription = "This is a test command"
		WithCommandDescription(commandDescription)(cmd)
		return commandDescription, cmd
	}

	t.Run("should set the command description", func(t *testing.T) {
		if commandDescription, cmd := runWithCommandDescriptionTest(); cmd.description != commandDescription {
			t.Errorf("Expected command description to be %s, got %s", commandDescription, cmd.description)
		}
	})
}

func TestWithCommandName(t *testing.T) {
	runWithCommandNameTest := func() (commandName string, cmd *MapCommand) {
		cmd = &MapCommand{}
		commandName = "test"
		WithCommandName(commandName)(cmd)
		return commandName, cmd
	}

	t.Run("should set the command name", func(t *testing.T) {
		if commandName, cmd := runWithCommandNameTest(); cmd.name != commandName {
			t.Errorf("Expected command name to be %s, got %s", commandName, cmd.name)
		}
	})
}

func TestWithEvictionConfig(t *testing.T) {
	runWithEvictionConfigTest := func() (ec *qec.QueryEvictionConfig, cmd *MapCommand) {
		cmd = &MapCommand{}
		ec = qec.NewQueryEvictionConfig()
		WithEvictionConfig(ec)(cmd)
		return ec, cmd
	}

	t.Run("should set the eviction configuration", func(t *testing.T) {
		if ec, cmd := runWithEvictionConfigTest(); ec != cmd.ec {
			t.Errorf("Expected eviction configuration to be %v, got %v", ec, cmd.ec)
		}
	})
}

func TestWithFetchLocations(t *testing.T) {
	runWithFetchLocationsTest := func() (fetchLocations FetchLocations, cmd *MapCommand) {
		cmd = &MapCommand{}
		fetchLocations = func(url string, config ...*qec.QueryEvictionConfig) (state *ms.MapsState, err error) {
			return state, err
		}
		WithFetchLocations(fetchLocations)(cmd)
		return fetchLocations, cmd
	}

	t.Run("should set the fetch locations function", func(t *testing.T) {
		fetchLocationsFunc, cmd := runWithFetchLocationsTest()
		utils.ExpectSameEntity(t, fetchLocationsFunc, cmd.fetchLocations, "FetchLocations")
	})
}

func TestWithListMarker(t *testing.T) {
	runWithListMarkerTest := func() (listMarker string, cmd *MapCommand) {
		cmd = &MapCommand{}
		listMarker = "-"

		WithListMarker(listMarker)(cmd)

		return listMarker, cmd
	}

	t.Run("should set the list marker", func(t *testing.T) {
		if listMarker, cmd := runWithListMarkerTest(); cmd.listMarker != listMarker {
			t.Errorf("Expected list marker to be %s, got %s", listMarker, cmd.listMarker)
		}
	})
}

func TestWithListTitle(t *testing.T) {
	runWithListTitleTest := func() (listTitle string, cmd *MapCommand) {
		cmd = &MapCommand{}
		listTitle = "Pokemon Maps:"

		WithListTitle(listTitle)(cmd)

		return listTitle, cmd
	}

	t.Run("should set the list title", func(t *testing.T) {
		if listTitle, cmd := runWithListTitleTest(); cmd.listTitle != listTitle {
			t.Errorf("Expected list title to be %s, got %s", listTitle, cmd.listTitle)
		}
	})
}

func TestWithNoMapsFoundErrorMessage(t *testing.T) {
	runWithNoMapsFoundErrorMessageTest := func() (noMapsFoundErrorMessage string, cmd *MapCommand) {
		cmd = &MapCommand{}
		noMapsFoundErrorMessage = "No maps found"

		WithNoMapsFoundErrorMessage(noMapsFoundErrorMessage)(cmd)

		return noMapsFoundErrorMessage, cmd
	}

	t.Run("should set the no maps found error message", func(t *testing.T) {
		if noMapsFoundErrorMessage, cmd := runWithNoMapsFoundErrorMessageTest(); cmd.noMapsFoundErrorMessage != noMapsFoundErrorMessage {
			t.Errorf("Expected no maps found error message to be %s, got %s", noMapsFoundErrorMessage, cmd.noMapsFoundErrorMessage)
		}
	})
}

func TestWithNoMoreMapsMessage(t *testing.T) {
	runWithNoMoreMapsMessageTest := func() (noMoreMapsMessage string, cmd *MapCommand) {
		cmd = &MapCommand{}
		noMoreMapsMessage = "No more maps"

		WithNoMoreMapsMessage(noMoreMapsMessage)(cmd)

		return noMoreMapsMessage, cmd
	}

	t.Run("should set the no more maps message", func(t *testing.T) {
		if noMoreMapsMessage, cmd := runWithNoMoreMapsMessageTest(); cmd.noMoreMapsMessage != noMoreMapsMessage {
			t.Errorf("Expected no more maps message to be %s, got %s", noMoreMapsMessage, cmd.noMoreMapsMessage)
		}
	})
}

func TestWithPaginationDirection(t *testing.T) {
	runWithPaginationDirectionTest := func() (paginationDirection string, cmd *MapCommand) {
		cmd = &MapCommand{}
		paginationDirection = pd.Next

		WithPaginationDirection(paginationDirection)(cmd)

		return paginationDirection, cmd
	}

	t.Run("should set the pagination direction", func(t *testing.T) {
		if paginationDirection, cmd := runWithPaginationDirectionTest(); cmd.paginationDirection != paginationDirection {
			t.Errorf("Expected pagination direction to be %s, got %s", paginationDirection, cmd.paginationDirection)
		}
	})
}

func TestWithState(t *testing.T) {
	runWithStateTest := func() (state *ms.MapsState, cmd *MapCommand) {
		cmd = &MapCommand{}
		state = ms.NewMapsState()

		WithState(state)(cmd)

		return state, cmd
	}

	t.Run("should set the state", func(t *testing.T) {
		if state, cmd := runWithStateTest(); state != cmd.state {
			t.Errorf("Expected state to be %v, got %v", state, cmd.state)
		}
	})
}
