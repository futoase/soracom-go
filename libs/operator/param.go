package operator

import (
	"github.com/codegangsta/cli"
	"strconv"
)

type RequestParams struct {
	XSoracomApiKey      string
	XSoracomToken       string
	OperatorId          string
	TokenTimeoutSeconds int
	CurrentPassword     string
	NewPassword         string
	Email               string
	Password            string
	VerifyToken         string
}

func (r *Request) SetParam(c *cli.Context) error {
	p := &r.Params
	if len(c.String("api-key")) != 0 {
		p.XSoracomApiKey = c.String("api-key")
	}

	if len(c.String("token")) != 0 {
		p.XSoracomToken = c.String("token")
	}

	if c.IsSet("operator-id") {
		p.OperatorId = c.String("operator-id")
	}

	if c.IsSet("token-timeout-seconds") {
		i, err := strconv.Atoi(c.String("token-timeout-seconds"))
		if err != nil {
			return err
		}
		p.TokenTimeoutSeconds = i
	}

	if len(c.String("current-password")) != 0 {
		p.CurrentPassword = c.String("current-password")
	}

	if len(c.String("new-password")) != 0 {
		p.NewPassword = c.String("new-password")
	}

	if len(c.String("email")) != 0 {
		p.Email = c.String("email")
	}

	if len(c.String("password")) != 0 {
		p.Password = c.String("password")
	}

	if len(c.String("verify-token")) != 0 {
		p.VerifyToken = c.String("verify-token")
	}

	return nil
}
