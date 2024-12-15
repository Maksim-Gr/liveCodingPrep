package pointers

import (
	"regexp"
	"strings"
)

func strRevere(s string, start, end int) string {
	runes := []rune(s)
	for start < end {
		runes[start], runes[end] = runes[end], runes[start]
		start++
		end--
	}
	return string(runes)
}

func reverseWords(sentence string) string {
	reg := regexp.MustCompile(`\s+`)
	sentence = reg.ReplaceAllString(sentence, " ")
	sentence = strings.TrimSpace(sentence)
	strLen := len(sentence) - 1

	sentence = strRevere(sentence, 0, len(sentence)-1)
	for start, end := 0, 0; end <= strLen; end++ {
		if end == strLen || sentence[end] == ' ' {
			var endIdx int
			if end == strLen {
				endIdx = end
			} else {
				endIdx = end - 1
			}
			sentence = strRevere(sentence, start, endIdx)
			start = end + 1
		}
	}
	return sentence
}
