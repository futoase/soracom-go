package auth

import (
  "github.com/codegangsta/cli"
)

func Exec(c *cli.Context) (*Response, string, error) {
  r := Request{}
  err := r.SetParam(c)
  if err != nil {
    return nil, "", err
  }

  resp, raw, err := r.Auth()
  if err != nil {
    return nil, "", err
  }

  return resp, raw, nil
}
