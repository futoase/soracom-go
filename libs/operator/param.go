package operator

import (
  "strconv"
  "github.com/codegangsta/cli"
)

func (r *Request) SetParam(c *cli.Context) error {
  if len(c.String("api-key")) != 0 {
    r.XSoracomApiKey = c.String("api-key")
  }

  if len(c.String("token")) != 0 {
    r.XSoracomToken = c.String("token")
  }

  if c.IsSet("operator-id") {
    r.OperatorId = c.String("operator-id")
  }

  if c.IsSet("token-timeout-seconds") {
    i, err := strconv.Atoi(c.String("token-timeout-seconds"))
    if err != nil {
      return err
    }
    r.TokenTimeoutSeconds = i
  }

  if len(c.String("current-password")) != 0 {
    r.CurrentPassword = c.String("current-password")
  }

  if len(c.String("new-password")) != 0 {
    r.NewPassword = c.String("new-password")
  }

  if len(c.String("email")) != 0 {
    r.Email = c.String("email")
  }

  if len(c.String("password")) != 0 {
    r.Password = c.String("password")
  }

  if len(c.String("verify-token")) != 0 {
    r.VerifyToken = c.String("verify-token")
  }

  return nil
}
