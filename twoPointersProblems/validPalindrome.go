package palindrome

// Check the valid palindrome using two pointers approach

// IsValidPalindrome method to check if string is valid palindrome
func IsValidPalindrome(inputString string) bool {
	left := 0
	right := len(inputString) - 1
	for left < right {
		if inputString[left] != inputString[right] {
			return false
		}
		left++
		right--
	}
	return true
}
