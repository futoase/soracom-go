package operator

import (
  "github.com/codegangsta/cli"
)

func ExecToken(c *cli.Context) (*Response, string, error) {
  r := Request{}
  err := r.SetParam(c)
  if err != nil {
    return nil, "", err
  }

  _, raw, err := r.OperatorToken()
  if err != nil {
    return nil, "", err
  }
  return nil, raw, nil
}

func ExecPassword(c *cli.Context) (*Response, string, error) {
  r := Request{}
  err := r.SetParam(c)
  if err != nil {
    return nil, "", err
  }

  _, raw, err := r.OperatorPassword()
  if err != nil {
    return nil, "", err
  }
  return nil, raw, nil
}
