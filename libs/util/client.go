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

func (c *HttpClient) NewRequest(method string, path string, body []byte) (*http.Request, error) {
  if body == nil {
    req, err := http.NewRequest(method, path, nil)
    if err != nil {
      return nil, err
    }
    return req, nil
  } else {
    req, err := http.NewRequest(method, path, bytes.NewReader(body))
    if err != nil {
      return nil, err
    }
    return req, nil
  }
}

func (c *HttpClient) Request(method string, path string, body []byte) (*http.Response, error) {
	client := &http.Client{}

	req, err := c.NewRequest(method, path, body)
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
