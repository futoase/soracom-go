package command

import (
	"github.com/codegangsta/cli"
)

func setFlag() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:   "api-key",
			Value:  "api-key",
			EnvVar: "SORACOM_API_KEY",
		},
		cli.StringFlag{
			Name:   "token",
			Value:  "token",
			EnvVar: "SORACOM_TOKEN",
		},
		cli.IntFlag{
			Name:  "operator-id",
			Usage: "operator-id",
		},
	}
}
