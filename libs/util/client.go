package util

import (
	"bytes"
	config "github.com/futoase/soracom-go/libs/config"
	"net/http"
)

type HttpClient struct {
	Path           string
	Body           []byte
	XSoracomApiKey string
	XSoracomToken  string
}

func ReadBody(body []byte) *bytes.Reader {
	if body != nil {
		return bytes.NewReader(body)
	} else {
		return nil
	}
}

func (c *HttpClient) Request(method string, path string, body []byte) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, path, ReadBody(body))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-Soracom-API-Key", c.XSoracomApiKey)
	req.Header.Add("X-Soracom-Token", c.XSoracomToken)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *HttpClient) Get() (*http.Response, error) {
	resp, err := c.Request("GET", config.API_ENDPOINT+c.Path, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *HttpClient) Post() (*http.Response, error) {
	resp, err := c.Request("POST", config.API_ENDPOINT+c.Path, c.Body)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
