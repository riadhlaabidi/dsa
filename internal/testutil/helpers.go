package testutil

import "testing"

func AssertEquals[T comparable](t *testing.T, want, got T) {
	t.Helper()
	if got != want {
		t.Fatalf("want: %v, got: %v", want, got)
	}
}
