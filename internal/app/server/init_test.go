package server

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetup(t *testing.T) {
	_, err := Setup()
	assert.Nil(t, err)
}
