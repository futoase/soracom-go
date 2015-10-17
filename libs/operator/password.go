package operator

import (
  "bytes"
  "encoding/json"
  "net/http"
  "net/http/httputil"
  "errors"
  "io/ioutil"
  config "github.com/futoase/soracom-go/libs/config"
)

func (r *Request) OperatorPassword() (*Response, string, error) {
  client := &http.Client{}

  rp := RequestPassword{r.CurrentPassword, r.NewPassword}

  reqBody, err := json.Marshal(rp)
  if err != nil {
    return nil, "", err
  }

  req, err := http.NewRequest("POST", config.API_ENDPOINT + "/operators/" + string(r.OperatorId) + "/password", bytes.NewReader(reqBody))
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
    err = errors.New("Invalid password")
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
