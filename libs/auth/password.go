package auth

import (
	"bytes"
	"encoding/json"
	"errors"
	util "github.com/futoase/soracom-go/libs/util"
	"io/ioutil"
	"net/http"
	"net/url"
)

func (r *Request) PassWordResetTokenTheIssue() (*Response, string, error) {
	ar := AuthRequest{r.Email}

	mJson, err := json.Marshal(ar)
	if err != nil {
		return nil, "", err
	}

	client := util.HttpClient{}
	client.Path = "/auth/password_reset_token/issue"
	client.Body = mJson

	resp, err := client.Post()
	if err != nil {
		return nil, "", err
	}

	if resp.StatusCode == http.StatusBadRequest {
		err = errors.New("Email is not valid.")
		return nil, "", err
	}

	if resp.StatusCode == http.StatusOK {
		return nil, "OK", nil
	}

	err = errors.New("Undefined error.")
	return nil, "", err
}

func (r *Request) PassWordResetTokenTheVerify() (*Response, string, error) {
	mJson, err := json.Marshal(r)
	if err != nil {
		return nil, "", err
	}

	contentReader := bytes.NewReader(mJson)
	resp, err := http.Post(config.API_ENDPOINT+"/auth/password_reset_token/verify", "application/json", contentReader)
	if err != nil {
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
