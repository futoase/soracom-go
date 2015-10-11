package auth

import (
  "net/http"
  "errors"
  "net/url"
  config "github.com/futoase/soracom/libs/config"
)

func (r *Request)PassWordResetTokenTheIssue() (*Response, string, error) {
  resp, err := http.PostForm(config.API_ENDPOINT + "/auth/password_reset_token/issue", url.Values{"email": {r.Email}})
  if err != nil {
    return nil, "", err
  }

  println(resp.StatusCode)

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

func (r *Request)PassWordResetTokenTheVerify() (*Response, string, error){
  println("password_reset_token verify")
  return nil, "", nil
}
