package operator

import (
	"encoding/json"
	"errors"
	config "github.com/futoase/soracom-go/libs/config"
	"io/ioutil"
	"net/http"
)

func (r *Request) OperatorInfo() (*Response, string, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", config.API_ENDPOINT+"/operators/"+string(r.OperatorId), nil)
	if err != nil {
		return nil, "", err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-Soracom-API-Key", r.XSoracomApiKey)
	req.Header.Add("X-Soracom-Token", r.XSoracomToken)

	resp, err := client.Do(req)
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
