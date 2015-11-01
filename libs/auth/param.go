package auth

import (
	"github.com/codegangsta/cli"
)

type RequestParams struct {
	Email               string
	Password            string
	TokenTimeoutSeconds int
	OneTimeToken        string
	VerifyToken         string
	NewPassword         string
}

func (r *Request) SetParam(c *cli.Context) error {
	p := &r.Params
	if len(c.String("email")) != 0 {
		p.Email = c.String("email")
	}

	if len(c.String("password")) != 0 {
		p.Password = c.String("password")
	}

	if c.IsSet("token-timeout-seconds") {
		p.TokenTimeoutSeconds = c.Int("token-timeout-seconds")
	} else {
		p.TokenTimeoutSeconds = 86400
	}

	if len(c.String("token")) != 0 {
		p.OneTimeToken = c.String("token")
	}

	if len(c.String("new-password")) != 0 {
		p.NewPassword = c.String("new-password")
	}

	if len(c.String("verify-token")) != 0 {
		p.VerifyToken = c.String("verify-token")
	}

	return nil
}
