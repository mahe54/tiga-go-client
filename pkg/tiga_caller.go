package tigaclient

import (
	"net/http"
	"time"
)

func (c *Caller) DoCall(req *http.Request) (*http.Response, error) {

	httpClient := &http.Client{Timeout: 30 * time.Second}
	response, err := httpClient.Do(req)
	if err != nil {
		return response, err
	}
	return response, err
}
