package testutil

import (
	"slices"
	"testing"
)

func AssertEquals[T comparable](t *testing.T, want, got T) {
	t.Helper()
	if got != want {
		t.Fatalf("want: %v, got: %v", want, got)
	}
}

func AssertNotNil(t *testing.T, got any) {
	t.Helper()
	if got == nil {
		t.Fatalf("want: non-nil, got: nil")
	}
}

func AssertSliceEquals[T comparable](t *testing.T, want, got []T) {
	t.Helper()
	if !slices.Equal(want, got) {
		t.Fatalf("want: %v, got: %v", want, got)
	}
}
