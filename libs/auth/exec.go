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

  switch {
  case c.IsSet("password_reset_token"):
    switch {
    case c.IsSet("issue"):
      _, raw, err := r.PassWordResetTokenTheIssue()
      if err != nil {
        return nil, "", err
      } else {
        return nil, raw, nil
      }
    case c.IsSet("verify"):
      r.PassWordResetTokenTheVerify()
      return nil, "", nil
    }
  case c.IsSet("auth"):
    resp, raw, err := r.Auth()
    if err != nil {
      return nil, "", err
    }

    return resp, raw, nil
  }

  return nil, "", nil
}
