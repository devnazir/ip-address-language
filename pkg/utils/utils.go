package utils

import (
	"fmt"
	"os"
	"path/filepath"
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
