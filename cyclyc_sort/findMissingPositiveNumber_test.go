package cyclyc_sort

import "testing"

func TestFindMissingPositiveNumber(t *testing.T) {
	var testCases = []struct {
		name     string
		input    []int
		received int
		want     int
	}{
		{"Should return 5", []int{1, 2, 3, 4}, 5, 5},
		{"Should return 2", []int{-1, 3, 5, 7, 1}, 2, 2},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			ans := smallestMissingPositiveInteger(tt.input)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
