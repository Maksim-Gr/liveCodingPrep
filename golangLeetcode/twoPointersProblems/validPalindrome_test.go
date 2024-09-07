package pointers

import "testing"

func TestIsValidPalindrome(t *testing.T) {
	var testCases = []struct {
		name  string
		input string
		want  bool
	}{
		{"Should fail", "ABC", false},
		{"Should pass", "A", true},
		{"Should pass", "RACECAR", true},
		{"Should fail", "RACEACAR", false},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			ans := IsValidPalindrome(tt.input)
			if ans != tt.want {
				t.Errorf("got %t, want %t", ans, tt.want)
			}
		})
	}
}
