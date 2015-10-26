package operator

import (
	"bytes"
	"encoding/json"
	"errors"
	config "github.com/futoase/soracom-go/libs/config"
	"net/http"
	"net/http/httputil"
)

func (r *Request) OperatorVerify() (*Response, string, error) {
	client := &http.Client{}

	rv := RequestVerify{r.VerifyToken}

	reqBody, err := json.Marshal(rv)
	if err != nil {
		return nil, "", err
	}

	req, err := http.NewRequest("POST", config.API_ENDPOINT+"/operators/verify", bytes.NewReader(reqBody))
	if err != nil {
		return nil, "", err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-Soracom-API-Key", r.XSoracomApiKey)
	req.Header.Add("X-Soracom-Token", r.XSoracomToken)

	body, err := httputil.DumpRequest(req, true)
	if err != nil {
		println(err)
	}
	println(string(body))

	resp, err := client.Do(req)
	if err != nil {
		return nil, "", err
	}

	switch resp.StatusCode {
	case http.StatusBadRequest:
		err = errors.New("Invalid token")
		return nil, "", err
	case http.StatusNotFound:
		err = errors.New("Token is expired")
		return nil, "", err
	case http.StatusOK:
		return nil, "OK", nil
	}

	println(resp.StatusCode)

	err = errors.New("Unknown error.")
	return nil, "", err
}
