package logger

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitialize(t *testing.T) {
	t.Run("valid logger initialization", func(t *testing.T) {
		err := Initialize("info")
		assert.NoError(t, err)
	})

	t.Run("invalid logger initialization", func(t *testing.T) {
		err := Initialize("none")
		assert.Error(t, err)
	})
}
