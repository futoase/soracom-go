package operator

import (
	"bytes"
	"encoding/json"
	"errors"
	config "github.com/futoase/soracom-go/libs/config"
	"net/http"
	"net/http/httputil"
)

func (r *Request) Operators() (*Response, string, error) {
	client := &http.Client{}

	ro := RequestOperators{r.Email, r.Password}

	reqBody, err := json.Marshal(ro)
	if err != nil {
		return nil, "", err
	}

	req, err := http.NewRequest("POST", config.API_ENDPOINT+"/operators", bytes.NewReader(reqBody))
	if err != nil {
		return nil, "", err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-Soracom-API-Key", r.XSoracomApiKey)
	req.Header.Add("X-Soracom-Token", r.XSoracomToken)

	requestByte, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		println(err)
	}
	println(string(requestByte))

	resp, err := client.Do(req)
	if err != nil {
		return nil, "", err
	}

	switch resp.StatusCode {
	case http.StatusCreated:
		return nil, "OK", nil
	case http.StatusBadRequest:
		err = errors.New("This email is already registered")
		return nil, "", err
	}

	err = errors.New("Unknown error.")
	return nil, "", err
}
