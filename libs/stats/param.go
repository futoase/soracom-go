package stats

import (
	"github.com/codegangsta/cli"
	"strconv"
)

type RequestParams struct {
	XSoracomApiKey string
	XSoracomToken  string
	Imsi           string
	From           int
	To             int
	Period         string
}

func (r *Request) SetParam(c *cli.Context) error {
	p := &r.Params
	if len(c.String("api-key")) != 0 {
		p.XSoracomApiKey = c.String("api-key")
	}

	if len(c.String("token")) != 0 {
		p.XSoracomToken = c.String("token")
	}

	if len(c.String("imsi")) != 0 {
		p.Imsi = c.String("imsi")
	}

	if len(c.String("from")) != 0 {
		fi, err := strconv.Atoi(c.String("from"))
		if err != nil {
			return err
		}
		p.From = fi
	}

	if len(c.String("to")) != 0 {
		ti, err := strconv.Atoi(c.String("to"))
		if err != nil {
			return err
		}
		p.To = ti
	}

	if len(c.String("period")) != 0 {
		p.Period = c.String("period")
	}

	return nil
}
