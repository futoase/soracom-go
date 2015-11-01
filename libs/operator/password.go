package operator

import (
	"encoding/json"
	"errors"
	util "github.com/futoase/soracom-go/libs/util"
	"net/http"
)

func (r *Request) OperatorPassword() (*Response, string, error) {
	rp := RequestPassword{r.CurrentPassword, r.NewPassword}

	mJson, err := json.Marshal(rp)
	if err != nil {
		return nil, "", err
	}

	client := util.HttpClient{}
	client.Path = "/operators/" + string(r.OperatorId) + "/password"
	client.Body = mJson
	client.XSoracomApiKey = r.XSoracomApiKey
	client.XSoracomToken = r.XSoracomToken

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
