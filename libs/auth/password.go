package auth

import (
	"encoding/json"
	"errors"
	util "github.com/futoase/soracom-go/libs/util"
	"net/http"
)

func (r *Request) PassWordResetTokenTheIssue() (*Response, string, error) {
	ir := IssueRequest{r.Email}

	mJson, err := json.Marshal(ir)
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
	vr := VerifyRequest{r.NewPassword, r.VerifyToken}

	mJson, err := json.Marshal(vr)
	if err != nil {
		return nil, "", err
	}

	client := util.HttpClient{}
	client.Path = "/auth/password_reset_token/verify"
	client.Body = mJson

	resp, err := client.Post()
	if err != nil {
		return nil, "", err
	}

	switch resp.StatusCode {
	case http.StatusBadRequest:
		err = errors.New("Invalid token.")
		return nil, "", err
	case http.StatusNotFound:
		err = errors.New("Token timeout.")
		return nil, "", err
	case http.StatusOK:
		return nil, "OK", nil
	}

	err = errors.New("Unknown Error.")
	return nil, "", err
}
