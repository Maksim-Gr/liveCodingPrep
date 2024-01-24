package pointers

import (
	"regexp"
	"strings"
)

func ReverseWordsInString(sentence string) string {
	trimmedSentence := trimStri
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
