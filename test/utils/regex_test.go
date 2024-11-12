package utils_test

import (
	"testing"

	"github.com/devnazir/gosh-script/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestFindSubShellArgs(t *testing.T) {
	text := "List of files: $(ls)"
	expectedArgs := []string{"$(ls)", "ls"}

	actualArgs := utils.FindSubShellArgs(text)
	assert.Equal(t, expectedArgs, actualArgs)
}
