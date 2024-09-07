package cyclyc_sort

import "testing"

func TestFindMissingNumber(t *testing.T) {
	var testCases = []struct {
		name     string
		input    []int
		received int
		want     int
	}{
		{"Should return 1", []int{0, 2, 5, 3, 4}, 1, 1},
		{"Should return 3", []int{0, 2, 5, 1, 4}, 1, 3},
		{"Should return 5", []int{0, 2, 1, 3, 4}, 5, 5},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			ans := findMissingNumber(tt.input)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
