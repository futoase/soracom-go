package auth

import (
  "strconv"
  "github.com/codegangsta/cli"
)

func (r *Request) SetParam(c *cli.Context) error {
  if len(c.String("email")) != 0 {
    r.Email = c.String("email")
  }

  if len(c.String("password")) != 0 {
    r.Password = c.String("password")
  }

  if c.IsSet("token-timeout-seconds") {
    i, err := strconv.Atoi(c.String("token-timeout-seconds"))
    if err != nil {
      return err
    }
    r.TokenTimeoutSeconds = i
  }

  if len(c.String("token")) != 0 {
    r.Token = c.String("token")
  }

  return nil
}
