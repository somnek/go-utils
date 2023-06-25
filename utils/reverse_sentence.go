package utils

import (
	"strings"
)

func ReverseSentence(s string) string {
	words := strings.Split(s, " ")
	reversedWords := make([]string, len(words))
	for i, j := 0, len(words)-1; i < len(words); i, j = i+1, j-1 {
		reversedWords[i] = words[j]
	}
	return strings.Join(reversedWords, " ")
}
