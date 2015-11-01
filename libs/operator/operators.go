package operator

import (
	"encoding/json"
	"errors"
	util "github.com/futoase/soracom-go/libs/util"
	"net/http"
)

func (r *Request) Operators() (*Response, string, error) {
	ro := RequestOperators{r.Email, r.Password}

	mJson, err := json.Marshal(ro)
	if err != nil {
		return nil, "", err
	}

	client := util.HttpClient{}
	client.Path = "/operators"
	client.Body = mJson
	client.XSoracomApiKey = r.XSoracomApiKey
	client.XSoracomToken = r.XSoracomToken

	resp, err := client.Post()
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
