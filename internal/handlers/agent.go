package handlers

import (
	"net/http"
)

func AgentRequest(url string) (*http.Response, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("POST", url, nil)
	req.Header.Set("Content-Type", "text/plain")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return resp, nil
}
