package operator

import (
	"encoding/json"
	"errors"
	config "github.com/futoase/soracom-go/libs/config"
	"io/ioutil"
	"net/http"
)

func (r *Request) OperatorSupportToken() (*Response, string, error) {
	client := &http.Client{}

	req, err := http.NewRequest("POST", config.API_ENDPOINT+"/operators/"+string(r.OperatorId)+"/support/token", nil)
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
