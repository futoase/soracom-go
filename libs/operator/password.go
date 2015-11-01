package operator

import (
	"encoding/json"
	"errors"
	util "github.com/futoase/soracom-go/libs/util"
	"net/http"
)

func (r *Request) OperatorPassword() (*Response, string, error) {
	p := &r.Params
	rp := RequestPassword{p.CurrentPassword, p.NewPassword}

	mJson, err := json.Marshal(rp)
	if err != nil {
		return nil, "", err
	}

	client := util.HttpClient{}
	client.Path = "/operators/" + string(p.OperatorId) + "/password"
	client.Body = mJson
	client.XSoracomApiKey = p.XSoracomApiKey
	client.XSoracomToken = p.XSoracomToken

	resp, err := client.Post()
	if err != nil {
		return nil, "", err
	}

	switch resp.StatusCode {
	case http.StatusBadRequest:
		err = errors.New("Invalid password")
		return nil, "", err
	case http.StatusOK:
		return nil, "OK", err
	}

	err = errors.New("Unknown Error.")
	return nil, "", err
}
