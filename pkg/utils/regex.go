package utils

import (
	"regexp"
)

func FindSubShellArgs(str string) []string {
	matcherArgs := regexp.MustCompile(`\$\((.*)\)`).FindStringSubmatch(str)
	return matcherArgs
}

func FindShellVars(str string) []string {
	matcherArgs := regexp.MustCompile(`\$\w+(\.\w+|\[\w+\])?`).FindAllString(str, -1)
	return matcherArgs
}

func IsShellVarMemberExpr(str string) (bool, string, []string, string) {
	rootVarRegex := regexp.MustCompile(`\$(\w+)`)
	memberExprRegex := regexp.MustCompile(`(\.\w+|\[\w+\])`)

	rootMatches := rootVarRegex.FindStringSubmatch(str)
	if len(rootMatches) < 2 {
		return false, "", nil, ""
	}
	rootShellVar := rootMatches[1]

	matchShellVar := memberExprRegex.MatchString(str)
	memberExprMatches := memberExprRegex.FindAllString(str, -1)

	var memberExprValue string
	if len(memberExprMatches) > 0 {
		if memberExprMatches[0][0] == '.' {
			memberExprValue = "object"
		} else {
			memberExprValue = "array"
		}
	}

	return matchShellVar, rootShellVar, memberExprMatches, memberExprValue
}
