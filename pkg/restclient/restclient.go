package restclient

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type restClient struct {
	client  http.Client
	baseURL string
}

func (r *restClient) DoGet(path string) ([]byte, error) {
	// Prepare GET request
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", r.baseURL, path), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-Caller-Scopes", "admin")

	// Do request
	resp, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	// Close body on finish
	defer resp.Body.Close()

	// Read response to bytes
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func NewRestClient(baseURL string) *restClient {
	return &restClient{
		client:  http.Client{},
		baseURL: baseURL,
	}
}
