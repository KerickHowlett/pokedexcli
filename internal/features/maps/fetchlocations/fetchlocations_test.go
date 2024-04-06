package fetchlocations

import (
	"fmt"
	"net/http"
	"testing"

	l "maps/location"
	s "query/state"
	f "testtools/fixtures"
	"testtools/utils"
)

func TestFetchLocations(t *testing.T) {
	setup := func(responseType string) (queryMapState *s.QueryState[l.Location], mockedFetch QueryFetchFunc, stdout *utils.PrintStorage) {
		queryMapState = s.NewQueryState[l.Location]()
		stdout = utils.NewPrintStorage()

		fetchedMock := func(responseType string) QueryFetchFunc {
			if responseType == "error" {
				return func(url string, queryState *s.QueryState[l.Location]) error {
					return fmt.Errorf("error with response: %s", http.StatusText(http.StatusInternalServerError))
				}
			}
			if responseType == "no-maps" {
				return func(url string, queryState *s.QueryState[l.Location]) error {
					queryState.Results = []l.Location{}
					return nil
				}
			}
			if responseType == "success" {
				return func(url string, queryState *s.QueryState[l.Location]) error {
					newLocation := l.NewLocation(l.WithName(f.StarterTown))
					queryState.Results = append(queryState.Results, *newLocation)
					return nil
				}
			}
			panic(fmt.Sprintf("Unexpected responseType used for test setup: %s", responseType))
		}(responseType)

		return queryMapState, fetchedMock, stdout
	}

	expectErrorMessage := func(err error, expectedErrorMessage string) {
		if err == nil {
			t.Error("Expected an error, but got nil.")
		}
		if actualErrorMessage := err.Error(); actualErrorMessage != expectedErrorMessage {
			t.Errorf("Expected error to be %q, but got %q.", expectedErrorMessage, actualErrorMessage)
		}
	}

	t.Run("should panic if the fetch function is nil.", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("Expected a panic to occur, but function proceeded with remaining operations.")
			}
		}()
		FetchLocations(nil, nil, nil)
	})

	t.Run("should return an error if the state is nil.", func(t *testing.T) {
		_, mockedFetch, _ := setup("success")
		err := FetchLocations(&f.APIEndpoint, nil, mockedFetch)
		expectErrorMessage(err, "a query state is required to fetch locations")
	})

	t.Run("should return an error if the url is nil.", func(t *testing.T) {
		state, mockedFetch, _ := setup("success")
		err := FetchLocations(nil, state, mockedFetch)
		expectErrorMessage(err, "no more maps to fetch")
	})

	t.Run("should return an error if the query fetch fails.", func(t *testing.T) {
		state, mockedFetch, _ := setup("error")
		err := FetchLocations(&f.APIEndpoint, state, mockedFetch)
		expectErrorMessage(err, http.StatusText(http.StatusInternalServerError))
	})

	t.Run("should return an error if no locations are found.", func(t *testing.T) {
		state, mockedFetch, _ := setup("no-maps")
		err := FetchLocations(&f.APIEndpoint, state, mockedFetch)
		expectErrorMessage(err, "no maps were found")
	})

	t.Run("should print all map location names stored in query state.", func(t *testing.T) {
		state, mockedFetch, stdout := setup("success")

		output, err := stdout.CaptureWithError(func() error {
			return FetchLocations(&f.APIEndpoint, state, mockedFetch)
		})

		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if expectedOutput := fmt.Sprintf("%s\n", f.StarterTown); output != expectedOutput {
			t.Errorf("Expected output to be %q, but got %q.", expectedOutput, output)
		}
	})
}