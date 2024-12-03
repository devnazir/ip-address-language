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