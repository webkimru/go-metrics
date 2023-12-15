package server

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRun(t *testing.T) {
	err := Setup()
	assert.Nil(t, err)
}
