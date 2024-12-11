package utils

import (
	"strings"
)

func RemoveDoubleQuotes(text string) (string, int) {
	if len(text) < 2 {
		return text, 0
	}

	doubleQuotesCount := strings.Count(text, "\"")
	text = strings.ReplaceAll(text, "\"", "")
	return text, doubleQuotesCount
}

func GetVariableName(text string) string {
	replacer := strings.NewReplacer("\"", "", "$", "", "{", "", "}", "")
	return replacer.Replace(text)
}
