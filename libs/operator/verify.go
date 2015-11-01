package operator

import (
	"encoding/json"
	"errors"
	util "github.com/futoase/soracom-go/libs/util"
	"net/http"
)

func (r *Request) OperatorVerify() (*Response, string, error) {
	p := &r.Params
	rv := RequestVerify{p.VerifyToken}

	mJson, err := json.Marshal(rv)
	if err != nil {
		return nil, "", err
	}

	client := util.HttpClient{}
	client.Path = "/operators/verify"
	client.Body = mJson
	client.XSoracomApiKey = p.XSoracomApiKey
	client.XSoracomToken = p.XSoracomToken

	resp, err := client.Post()
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
