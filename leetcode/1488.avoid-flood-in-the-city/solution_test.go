package avoidflood

import (
	"dsa/internal/testutil"
	"testing"
)

func TestAvoidFlood(t *testing.T) {
	tests := []struct {
		name  string
		rains []int
		want  []int
	}{
		{"Should not flood when raining all over distinct lakes", []int{1, 2, 3, 4}, []int{-1, -1, -1, -1}},
		{"Should not flood when dry chances are equal inside the rain interval", []int{1, 2, 0, 0, 2, 1}, []int{-1, -1, 2, 1, -1, -1}},
		{"Should flood when there is not enough dry days", []int{1, 2, 0, 1, 2}, []int{}},
		{"Should not flood when there is extra dry days inside the rain interval", []int{69, 0, 0, 0, 69}, []int{-1, 69, 1, 1, -1}},
		{"Should flood when dry days are out of rain interval", []int{0, 1, 1}, []int{}},
		{"Should not flood when there is enough dry days between each lake rain interval", []int{1, 0, 2, 0, 2, 1}, []int{-1, 1, -1, 2, -1, -1}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			testutil.AssertSliceEquals(t, test.want, avoidFlood(test.rains))
		})
	}
}
