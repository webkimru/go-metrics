package server

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetup(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	_, err := Setup(ctx)
	assert.Nil(t, err)
	cancel()
}
