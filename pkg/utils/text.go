package utils

import (
	"regexp"
	"strings"
)

func ExtractNonVarAndVars(text string) (string, string) {
	// Match all text that does not start with `$`
	nonVarRe := regexp.MustCompile(`(?:^|[^$])\b([A-Za-z\s]+)\b(?:[^$]|$)`)
	// Match all variables that start with `$`
	varRe := regexp.MustCompile(`\$\w+`)

	// Find all non-variable parts
	nonVarMatches := nonVarRe.FindAllString(text, -1)
	nonVarText := ""
	for _, match := range nonVarMatches {
		nonVarText += match
	}

	// Find all variable parts
	varMatches := varRe.FindAllString(text, -1)
	varText := ""
	for _, match := range varMatches {
		varText += match + " "
	}

	return nonVarText, varText
}

func RemoveDoubleQuotes(text string) (string, int) {
	if len(text) < 2 {
		return text, 0
	}

	text = strings.ReplaceAll(text, "\"", "")
	doubleQuotesCount := strings.Count(text, "\"")
	return text, doubleQuotesCount
}
