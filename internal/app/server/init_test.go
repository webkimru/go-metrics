package server

import (
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
	"os"
	"testing"
)

func TestSetup(t *testing.T) {
	t.Run("valid host address", func(t *testing.T) {
		os.Setenv("ADDRESS", "localhost:8080")
		os.Setenv("STORE_INTERVAL", "1")
		os.Setenv("RESTORE", "1")
		os.Setenv("KEY", "123")

		ctx, cancel := context.WithCancel(context.Background())
		_, err := Setup(ctx)
		assert.Nil(t, err)
		cancel()
	})
}
