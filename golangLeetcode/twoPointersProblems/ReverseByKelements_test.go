package pointers

import (
	"reflect"
	"testing"
)

func TestReverseArray(t *testing.T) {
	var testCases = []struct {
		name       string
		inputArray []int
		k          int
		want       []int
	}{
		{"should reverse array", []int{1, 2, 3, 4, 5}, 3, []int{3, 4, 5, 1, 2}},
		{"should reverse array by negative k", []int{1, 2, 3, 4, 5}, -2, []int{3, 4, 5, 1, 2}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ans := RotateArryByKlements(tc.inputArray, tc.k)
			if !reflect.DeepEqual(tc.want, ans) {
				t.Fatalf("expected: %v, got: %v", tc.want, ans)
			}
		})
	}
}
