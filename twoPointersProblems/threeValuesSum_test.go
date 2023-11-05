package pointers

import "testing"

func TestThreeValuesSum(t *testing.T) {

	var testCases = []struct {
		name  string
		input []int
		sum   int
		want  bool
	}{
		{"Should return false", []int{3, 7, 1, 2, 8, 4, 5}, 34, false},
		{"Should return false", []int{3, 7, 1, 2, 8, 4, 5}, 34, false},
		{"Should return false", []int{3, 7, 1, 2, 8, 4, 5}, 34, false},
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
