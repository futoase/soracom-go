package operator

import (
	"encoding/json"
	"errors"
	util "github.com/futoase/soracom-go/libs/util"
	"io/ioutil"
	"net/http"
)

func (r *Request) OperatorInfo() (*Response, string, error) {
	client := util.HttpClient{}
	client.Path = "/operators/" + string(r.OperatorId)
	client.XSoracomApiKey = r.XSoracomApiKey
	client.XSoracomToken = r.XSoracomToken

	resp, err := client.Get()
	if err != nil {
		return nil, "", err
	}

	if resp.StatusCode == http.StatusBadRequest {
		err = errors.New("Invalid Operator Id.")
		return nil, "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()

	or := Response{}

	err = json.Unmarshal(body, &or)
	if err != nil {
		return nil, "", err
	}

	return &or, string(body), nil
}
