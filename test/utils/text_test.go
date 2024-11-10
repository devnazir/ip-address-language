package utils_test

import (
	"testing"

	"github.com/devnazir/gosh-script/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestRemoveDoubleQuotes(t *testing.T) {
	text := "\"Hello, World!\""
	expectedText := "Hello, World!"
	expectedDoubleQuotesCount := 2

	actualText, actualDoubleQuotesCount := utils.RemoveDoubleQuotes(text)

	assert.Equal(t, expectedText, actualText)
	assert.Equal(t, expectedDoubleQuotesCount, actualDoubleQuotesCount)

	t.Run("Empty text", func(t *testing.T) {
		text := ""
		expectedText := ""
		expectedDoubleQuotesCount := 0

		actualText, actualDoubleQuotesCount := utils.RemoveDoubleQuotes(text)

		assert.Equal(t, expectedText, actualText)
		assert.Equal(t, expectedDoubleQuotesCount, actualDoubleQuotesCount)
	})
}
