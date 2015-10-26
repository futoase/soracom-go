package auth

import (
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
		r.TokenTimeoutSeconds = c.Int("token-timeout-seconds")
	} else {
		r.TokenTimeoutSeconds = 86400
	}

	if len(c.String("token")) != 0 {
		r.OneTimeToken = c.String("token")
	}

	return nil
}
