package auth

import (
  "encoding/json"
  "net/http"
  "errors"
  "net/url"
  "bytes"
  "io/ioutil"
  config "github.com/futoase/soracom-go/libs/config"
)

func (r *Request)PassWordResetTokenTheIssue() (*Response, string, error) {
  resp, err := http.PostForm(config.API_ENDPOINT + "/auth/password_reset_token/issue", url.Values{"email": {r.Email}})
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

func (r *Request)PassWordResetTokenTheVerify() (*Response, string, error){
  mJson, err := json.Marshal(r)
  if err != nil {
    return nil, "", err
  }

  contentReader := bytes.NewReader(mJson)
  resp, err := http.Post(config.API_ENDPOINT + "/auth/password_reset_token/verify", "application/json", contentReader)
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
