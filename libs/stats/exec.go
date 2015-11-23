package stats

import (
	"github.com/codegangsta/cli"
)

func ExecSubScriberStatus(c *cli.Context) (*SubScriberStatusResponse, string, error) {
	r := Request{}
	err := r.SetParam(c)
	if err != nil {
		return nil, "", err
	}

	_, raw, err := r.SubscriberStatus()
	if err != nil {
		return nil, "", err
	}

	return nil, raw, nil
}
