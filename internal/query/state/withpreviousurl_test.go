package querystate

import (
	"testing"

	f "testtools/fixtures"
)

func TestWithPreviousURL(t *testing.T) {
	var nilString *string = nil

	runWithPreviousURLTest := func(url *string, defaultPreviousURL ...*string) *queryTestState {
		state := &queryTestState{}
		if len(defaultPreviousURL) > 0 && defaultPreviousURL[0] != nil {
			state.PreviousURL = defaultPreviousURL[0]
		}

		WithPreviousURL[result](url)(state)

		return state
	}

	t.Run("should set QueryState[TResult].PreviousURL with the provided URL.", func(t *testing.T) {
		if state := runWithPreviousURLTest(&f.APIEndpoint); *state.PreviousURL != f.APIEndpoint {
			t.Errorf("Expected QueryState[TResult].PreviousURL to be %q, but got %q", f.APIEndpoint, *state.PreviousURL)
		}
	})

	t.Run("should set QueryState[TResult].PreviousURL with nil URL.", func(t *testing.T) {
		if state := runWithPreviousURLTest(nilString, &f.APIEndpoint); state.PreviousURL != nilString {
			t.Errorf("Expected QueryState[TResult].PreviousURL to be nil, but got non-nil")
		}
	})
}
