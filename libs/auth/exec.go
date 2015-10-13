package auth

import (
  "github.com/codegangsta/cli"
)

func ExecPasswordResetTheIssue(c *cli.Context) (*Response, string, error) {
  r := Request{}
  err := r.SetParam(c)
  if err != nil {
    return nil, "", err
  }

  _, raw, err := r.PassWordResetTokenTheIssue()
  if err != nil {
    return nil, "", err
  }
  return nil, raw, nil
}

func ExecPasswordResetTheVerify(c *cli.Context) (*Response, string, error) {
  r := Request{}
  err := r.SetParam(c)
  if err != nil {
    return nil, "", err
  }

  _, raw, err := r.PassWordResetTokenTheVerify()
  if err != nil {
    return nil, "", err
  }
  return nil, raw, nil
}

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
