package auth

import (
  "strconv"
  "github.com/codegangsta/cli"
)

func (r *Request) SetParam(c *cli.Context) error {
  if c.IsSet("email") {
    r.Email = c.String("email")
  }

  if c.IsSet("password") {
    r.Password = c.String("password")
  }

  if c.IsSet("token-timeout-seconds") {
    i, err := strconv.Atoi(c.String("token-timeout-seconds"))
    if err != nil {
      return err
    }
    r.TokenTimeoutSeconds = i
  }

  return nil
}
