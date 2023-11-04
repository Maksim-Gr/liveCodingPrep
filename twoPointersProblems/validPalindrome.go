package main

import "fmt"

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
	str := []string{"ABC", "A", "ABC"}
	for i, s := range str {
		fmt.Printf("Test Case # %d\n", i+1)
		fmt.Printf("\nIs it a palindrome?.....%v\n", IsValidPalindrome(s))
	}

}
