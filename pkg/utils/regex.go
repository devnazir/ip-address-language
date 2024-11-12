package utils

import "regexp"

func FindSubShellArgs(str string) []string {
	matcherArgs := regexp.MustCompile(`\$\((.*)\)`).FindStringSubmatch(str)
	return matcherArgs
}
