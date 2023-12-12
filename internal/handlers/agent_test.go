package handlers

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestAgentRequestPositive(t *testing.T) {
	t.Run("", func(t *testing.T) {
		resp, err := AgentRequest("http://localhost:8080/update/counter/someMetric/123")
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})
}

func TestAgentRequestNegative(t *testing.T) {
	t.Run("", func(t *testing.T) {
		resp, err := AgentRequest("http://abc")
		resp.Body.Close()
		assert.Error(t, err)
		assert.Nil(t, resp)
	})
}

func TestAgentRequestNegativeEmptyUrl(t *testing.T) {
	t.Run("", func(t *testing.T) {
		resp, err := AgentRequest("")
		resp.Body.Close()
		assert.Error(t, err)
		assert.Nil(t, resp)
	})
}

//func TestAgentRequestNoError(t *testing.T) {
//	t.Run("", func(t *testing.T) {
//		resp, err := AgentRequest("http://localhost:8080/update/counter/someMetric/123")
//		assert.NoError(t, err)
//		assert.Equal(t, http.StatusOK, resp.StatusCode)
//	})
//}
