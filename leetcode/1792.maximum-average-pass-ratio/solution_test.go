package main

import (
	"math"
	"testing"
)

func TestMaxAverageRatio(t *testing.T) {
	classes := [][]int{
		{1, 2},
		{3, 5},
		{2, 2},
	}
	extraStudents := 2

	want := 0.78333
	got := math.Round(maxAverageRatio(classes, extraStudents)*1e5) / 1e5

	if got != want {
		t.Errorf("Max average ratio, want = %f, got = %f", want, got)
	}
}
