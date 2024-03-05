package httpcl

import (
	"errors"
	"net/http"
)

func (c *httpClient) do(method string, url string, headers http.Header, body interface{}) (*http.Response, error) {
	client := http.Client{}

	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, errors.New("failed to create a new request")
	}

	// headers is of type Header map[string][]string
	for header, value := range headers {
		if len(value) > 0{
			request.Header.Set(header, value[0])
		}
	}

	return client.Do(request)
}
