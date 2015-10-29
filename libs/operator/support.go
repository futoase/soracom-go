package operator

import (
	"encoding/json"
	"errors"
	util "github.com/futoase/soracom-go/libs/util"
	"io/ioutil"
	"net/http"
)

func (r *Request) OperatorSupportToken() (*Response, string, error) {
	client := util.HttpClient{}
	client.Path = "/operators/" + string(r.OperatorId) + "/support/token"
	client.XSoracomApiKey = r.XSoracomApiKey
	client.XSoracomToken = r.XSoracomToken

	resp, err := client.Post()

	if err != nil {
		return nil, "", err
	}

	switch resp.StatusCode {
	case http.StatusBadRequest:
		err = errors.New("Invald OperatorId")
		return nil, "", err
	case http.StatusForbidden:
		err = errors.New("Invalid token for SORACOM")
		return nil, "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()

	ar := Response{}

	err = json.Unmarshal(body, &ar)
	if err != nil {
		return nil, "", err
	}

	return &ar, string(body), nil
}
