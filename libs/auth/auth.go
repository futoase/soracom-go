package auth

import (
  "encoding/json"
  "net/http"
  "errors"
  "bytes"
  "io/ioutil"
  config "github.com/futoase/soracom/libs/config"
)

func (r *Request) Auth() (*Response, string, error) {
  mJson, err := json.Marshal(r)
  if err != nil {
    return nil, "", err
  }

  contentReader := bytes.NewReader(mJson)
  resp, err := http.Post(config.API_ENDPOINT + "/auth", "application/json", contentReader)
  if err != nil {
    return nil, "", err
  }

  if resp.StatusCode == http.StatusForbidden {
    err = errors.New("Email or password is not valid.")
    return nil, "", err
  }
  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return nil, "", err
  }
  ar := Response{}

  err = json.Unmarshal(body, &ar)
  if err != nil {
    return nil, "", err
  }

  return &ar, string(body), nil
}
