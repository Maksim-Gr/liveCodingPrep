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
