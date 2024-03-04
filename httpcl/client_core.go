package httpcl

import (
	"errors"
	"net/http"
)

func (c *httpClient) do(method string, url string, headers http.Header, body interface{}) (*http.Response, error) {
	client := http.Client{}

	request, err := http.NewRequest(method, url, body: nil)
	if err != nil {
		return nil, errors.New("failed to create a new")
	}

	return client.Do(request)
}