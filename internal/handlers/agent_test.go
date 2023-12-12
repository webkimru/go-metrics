package handlers

//func TestAgentRequestPositive(t *testing.T) {
//	t.Run("positive test", func(t *testing.T) {
//		resp, err := AgentRequest("http://localhost:8080/update/counter/someMetric/123")
//		require.NoError(t, err)
//		defer resp.Body.Close()
//		assert.Equal(t, http.StatusOK, resp.StatusCode)
//	})
//}

//
//func TestAgentRequestNegative(t *testing.T) {
//	t.Run("", func(t *testing.T) {
//		resp, err := AgentRequest("http://abc")
//		require.NoError(t, err)
//		defer resp.Body.Close()
//		assert.Nil(t, resp)
//	})
//}
//
//func TestAgentRequestNegativeEmptyUrl(t *testing.T) {
//	t.Run("", func(t *testing.T) {
//		resp, err := AgentRequest("")
//		require.NoError(t, err)
//		defer resp.Body.Close()
//		assert.Nil(t, resp)
//	})
//}

//func TestAgentRequestNoError(t *testing.T) {
//	t.Run("", func(t *testing.T) {
//		resp, err := AgentRequest("http://localhost:8080/update/counter/someMetric/123")
//		assert.NoError(t, err)
//		assert.Equal(t, http.StatusOK, resp.StatusCode)
//	})
//}
