package query_fetch

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	qs "query/state"
	f "test_tools/fixtures"
)

type result struct {
	Name string `json:"name"`
}

func TestParseQueryFetchResponse(t *testing.T) {
	t.Run("should parse the response body correctly", func(t *testing.T) {
		body := []byte(fmt.Sprintf(`{
			"NextURL": null,
			"PreviousURL": null,
			"Results": [{"name": "%s"}]
		}`, f.StarterTown))
		state := qs.NewQueryState[result]()
		expectedState := qs.NewQueryState(
			qs.WithResult(result{Name: f.StarterTown}),
		)

		if parseQueryFetchResponse(body, state); !jsonEqual(state, expectedState) {
			t.Errorf("Expected QueryState[TResult] to be %v, but instead received %v", expectedState, state)
		}
	})

	t.Run("should return an error if the response body cannot be parsed", func(t *testing.T) {
		err := parseQueryFetchResponse(nil, qs.NewQueryState[result]())

		if err == nil {
			t.Errorf("Expected error to be returned, but got nil")
		}

		const expectedErrorMessage = "error with parsing payload"
		if !strings.Contains(err.Error(), expectedErrorMessage) {
			t.Errorf("Expected error message to be '%s', but instead received '%s'", expectedErrorMessage, err.Error())
		}
	})
}

// Helper function to compare JSON objects
func jsonEqual(a, b interface{}) bool {
	ajson, _ := json.Marshal(a)
	bjson, _ := json.Marshal(b)
	return string(ajson) == string(bjson)
}
