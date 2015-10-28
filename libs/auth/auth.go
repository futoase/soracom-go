package auth

import (
	"encoding/json"
	"errors"
	util "github.com/futoase/soracom-go/libs/util"
	"io/ioutil"
	"net/http"
)

func (r *Request) Auth() (*Response, string, error) {
	mJson, err := json.Marshal(r)
	if err != nil {
		return nil, "", err
	}

	client := util.HttpClient{}
	client.Path = "/auth"
	client.Body = mJson

	resp, err := client.Post()
	if err != nil {
		return nil, "", err
	}

	if resp.StatusCode == http.StatusForbidden {
		err = errors.New("Email or password is not valid.")
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
