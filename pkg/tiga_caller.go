package tigaclient

import (
	"net/http"
	"time"
)

func (c *Caller) DoCall(req *http.Request) (*http.Response, error) {

	// req, err := http.NewRequest(string(verb), url, bytes.NewBuffer(payload))
	// if err != nil {
	// 	fmt.Print(err.Error())
	// }
	// for key, val := range headers {
	// 	req.Header.Add(key, val)
	// }

	httpClient := &http.Client{Timeout: 10 * time.Second}
	response, err := httpClient.Do(req)
	if err != nil {
		return response, err
	}
	// defer resp.Body.Close()
	// response, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Print(err.Error())
	// }

	return response, err
}

// old
// func (c *Client) doRequest(req *http.Request) ([]byte, error) {

// 	res, err := c.hTTPClient.Do(req)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer res.Body.Close()

// 	body, err := ioutil.ReadAll(res.Body)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return body, err
// }
