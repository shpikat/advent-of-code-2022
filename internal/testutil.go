package internal

import (
	"os"
	"testing"
)

func ReadInput(t testing.TB, filename string) string {
	t.Helper()
	input, err := os.ReadFile(filename)
	if err != nil {
		t.Fatalf("Error reading file %s: %v", filename, err)
	}
	return string(input)
}
