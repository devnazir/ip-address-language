package utils_test

import (
	"testing"

	"github.com/devnazir/ip-address-language/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestFindSubShellArgs(t *testing.T) {
	text := "List of files: $(ls)"
	expectedArgs := []string{"$(ls)", "ls"}

	actualArgs := utils.FindSubShellArgs(text)
	assert.Equal(t, expectedArgs, actualArgs)
}

func TestFindShellVars(t *testing.T) {
	text := "Hello, $name! You have $count apples."
	expectedVars := []string{"$name", "$count"}

	actualVars := utils.FindShellVars(text)
	assert.Equal(t, expectedVars, actualVars)

	t.Run("Shell Var with dot notation", func(t *testing.T) {
		text := "Hello, $user.name! You have $user.count apples."
		expectedVars := []string{"$user.name", "$user.count"}

		actualVars := utils.FindShellVars(text)
		assert.Equal(t, expectedVars, actualVars)
	})

	t.Run("Shell Var with array notation", func(t *testing.T) {
		text := "Hello, $user[0]! You have $user[1] apples."
		expectedVars := []string{"$user[0]", "$user[1]"}

		actualVars := utils.FindShellVars(text)
		assert.Equal(t, expectedVars, actualVars)
	})
}
