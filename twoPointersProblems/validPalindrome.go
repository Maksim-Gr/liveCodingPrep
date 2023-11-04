package main

// Check the valid palindrome using two pointers approach

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

func main() {

}
