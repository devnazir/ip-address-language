package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func FindDirByFilename(root, filename string) (string, error) {
	var dir string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.Name() == filename {
			dir = filepath.Dir(path)
			return filepath.SkipDir
		}
		return nil
	})

	if err != nil {
		return "", err
	}
	if dir == "" {
		return "", fmt.Errorf("file not found: %s", filename)
	}

	return dir, nil
}

func IsComment(ch byte) bool {
	return ch == '/' || ch == '*'
}

func IsValidSyntax(ch byte) bool {
	return (ch >= '0' && ch <= '9') || ch == '.'
}

func IsNumeric(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func IsAlpha(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func TranslateTokenValue(char string) string {
	charArr := strings.Split(char, ".")
	word := ""

	if char == "" {
		return ""
	}

	if len(charArr) > 1 {
		for _, c := range charArr {
			num, err := strconv.Atoi(c)
			if err != nil {
				panic(err)
			}
			word += string(num)
		}
	}

	if len(charArr) == 1 {
		num, err := strconv.Atoi(char)
		if err != nil {
			panic(err)
		}

		word = string(num)
	}

	return strings.TrimSpace(word)
}
