package repl

import (
	"fmt"
	"testing"

	m "internal/tests/mocks"
	u "internal/tests/utils"
)

func TestWithScanner(t *testing.T) {
	fmt.Println("WithScanner should set the Scanner field of the REPL struct")

	scanner := m.NewScannerMock()
	repl := setupREPLOptionTest(WithScanner(scanner))

	u.ExpectSameEntity(t, repl.Scanner, scanner, "Scanner")
}