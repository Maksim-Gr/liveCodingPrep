package pointers

import (
	"regexp"
	"strings"
)

func ReverseWordsInString(sentence string) string {
	trimmedSentence := trimString(sentence)
	sentenceBytes := []byte(trimmedSentence)
	strLen := len(sentenceBytes)
	sentenceBytes = strRev(sentenceBytes, 0, strLen-1)
	start, end := 0, 0

	for start < strLen {

		for end < strLen && sentenceBytes[end] != ' ' {
			end += 1
		}

		// Let's call our helper function to reverse the word in-place.
		strRev(sentenceBytes, start, end-1)
		start = end + 1
		end += 1
	}
	return string(sentenceBytes)
}

func trimString(str string) string {
	str = strings.TrimSpace(str)

	regex := regexp.MustCompile("\\s+")
	str = regex.ReplaceAllString(str, " ")

	return str
}

func strRev(str []byte, startRev int, endRev int) []byte {
	for startRev < endRev {
		temp := str[startRev]
		str[startRev] = str[endRev]
		str[endRev] = temp

		startRev += 1
		endRev -= 1
	}
	return str
}
