package pointers

import "testing"

func TestFindSumOfThree(t *testing.T) {

	var testCases = []struct {
		name  string
		input []int
		sum   int
		want  bool
	}{
		{"Should return true, sum is exists", []int{3, 7, 1, 2, 8, 4, 5}, 10, true},
		{"Should return false, sum is not exists", []int{-1, 2, 1, -4, 5, -3}, 34, false},
		{"Should return false,sum is not exists", []int{3, 7, 1, 2, 8, 4, 5}, 34, false},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			ans := FindSumOfThree(tt.input, tt.sum)
			if ans != tt.want {
				t.Errorf("got %t, want %t", ans, tt.want)
			}
		})
	}
}
