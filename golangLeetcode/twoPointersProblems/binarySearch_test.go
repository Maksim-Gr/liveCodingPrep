package pointers

import "testing"

func TestBinarySearch(t *testing.T) {
	var testsCases = []struct {
		name     string
		inputArr []int
		target   int
		want     int
	}{
		{"Should return 5", []int{1, 2, 3, 4, 5, 6, 7}, 4, 3},
		{"Should return 1", []int{1, 2, 3, 4, 5, 6, 7}, 1, 0},
	}
	for _, tc := range testsCases {
		t.Run(tc.name, func(t *testing.T) {
			ans := BinarySearch(tc.inputArr, tc.target)
			if ans != tc.want {
				t.Errorf("got: %v, want: %v", ans, tc.want)
			}
		})
	}
}
