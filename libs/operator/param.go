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

  return nil
}
