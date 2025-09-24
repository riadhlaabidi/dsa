package versionnumbers

import (
	"dsa/internal/testutil"
	"testing"
)

func TestCompareVersions(t *testing.T) {
	var tests = []struct {
		name     string
		version1 string
		version2 string
		want     int
	}{
		{"Bigger revision number", "1.2", "1.10", -1},
		{"Same version with leading zeroes", "1.002.03.0001", "1.0002.0003.01", 0},
		{"Less revisions", "1.0", "1.0.0.0", 0},
		{"Bigger major version", "2.0", "1.9", 1},
		{"Major version only", "1", "3", -1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			testutil.AssertEquals(t, test.want, compareVersions(test.version1, test.version2))
		})
	}
}
