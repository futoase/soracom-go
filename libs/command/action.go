package command

import (
	"github.com/codegangsta/cli"
	auth "github.com/futoase/soracom-go/libs/auth"
	operator "github.com/futoase/soracom-go/libs/operator"
	"log"
)

func SetAction(a *cli.App) {
	a.Commands = []cli.Command{
		{
			Name:  "auth",
			Usage: "authorize",
			Action: func(c *cli.Context) {
				_, raw, err := auth.Exec(c)
				if err != nil {
					log.Fatal(err)
					return
				}
				println(string(raw))
				return
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:   "email, e",
					Value:  "email",
					EnvVar: "SORACOM_EMAIL",
				},
				cli.StringFlag{
					Name:   "password, p",
					Value:  "password",
					EnvVar: "SORACOM_PASSWORD",
				},
				cli.IntFlag{
					Name:  "token-timeout-seconds, t",
					Usage: "Timeout seconds of token",
					Value: 86400,
				},
			},
		},
		{
			Name:  "password-reset",
			Usage: "password reset",
			Subcommands: []cli.Command{
				{
					Name:  "issue",
					Usage: "get token of password reset from email",
          Flags: []cli.Flag{
				    cli.StringFlag{
				    	Name:   "email, e",
				    	Value:  "email",
				    	EnvVar: "SORACOM_EMAIL",
				    },
          },
					Action: func(c *cli.Context) {
						_, raw, err := auth.ExecPasswordResetTheIssue(c)
						if err != nil {
							log.Fatal(err)
							return
						}
						println(string(raw))
						return
					},
				},
				{
					Name:  "verify",
					Usage: "verify token of password reset",
          Flags: []cli.Flag{
				    cli.StringFlag{
				    	Name: "new-password, p",
				    	Value: "new-password",
				    },
            cli.StringFlag{
              Name: "verify-token, t",
              Value: "verify-token",
            },
          },
					Action: func(c *cli.Context) {
						_, raw, err := auth.ExecPasswordResetTheVerify(c)
						if err != nil {
							log.Fatal(err)
							return
						}
						println(string(raw))
						return
					},
				},
			},
		},
		{
			Name:  "operator",
			Usage: "operator",
			Subcommands: []cli.Command{
				{
					Name:  "token",
					Usage: "Update token of operator",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:   "api-key, a",
							Value:  "api-key",
							Usage:  "api key for SORACOM",
							EnvVar: "SORACOM_API_KEY",
						},
						cli.StringFlag{
							Name:   "token, t",
							Value:  "token",
							Usage:  "token for SORACOM",
							EnvVar: "SORACOM_TOKEN",
						},
						cli.IntFlag{
							Name:  "token-timeout-seconds, timeout",
							Usage: "Timeout seconds of token (default 86400)",
							Value: 86400,
						},
						cli.StringFlag{
							Name:  "operator-id, i",
							Value: "operator-id",
							Usage: "Set operator id",
						},
					},
					Action: func(c *cli.Context) {
						_, raw, err := operator.ExecToken(c)
						if err != nil {
							log.Fatal(err)
							return
						}
						println(string(raw))
						return
					},
				},
				{
					Name:  "password",
					Usage: "Update password",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:   "api-key, a",
							Value:  "api-key",
							Usage:  "api key for SORACOM",
							EnvVar: "SORACOM_API_KEY",
						},
						cli.StringFlag{
							Name:   "token, t",
							Value:  "token",
							Usage:  "token for SORACOM",
							EnvVar: "SORACOM_TOKEN",
						},
						cli.StringFlag{
							Name:  "operator-id, i",
							Value: "operator-id",
							Usage: "Set operator id",
						},
						cli.StringFlag{
							Name:  "current-password, c",
							Value: "current-password",
							Usage: "current password for SORACOM",
						},
						cli.StringFlag{
							Name:  "new-password, n",
							Value: "new-password",
							Usage: "new password for SORACOM",
						},
					},
					Action: func(c *cli.Context) {
						_, raw, err := operator.ExecPassword(c)
						if err != nil {
							log.Fatal(err)
							return
						}
						println(string(raw))
						return
					},
				},
				{
					Name:  "support_token",
					Usage: "get of token for support",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:   "api-key, a",
							Value:  "api-key",
							Usage:  "api key for SORACOM",
							EnvVar: "SORACOM_API_KEY",
						},
						cli.StringFlag{
							Name:   "token, t",
							Value:  "token",
							Usage:  "token for SORACOM",
							EnvVar: "SORACOM_TOKEN",
						},
						cli.StringFlag{
							Name:  "operator-id, i",
							Value: "operator-id",
							Usage: "Set operator id",
						},
					},
					Action: func(c *cli.Context) {
						_, raw, err := operator.ExecSupportToken(c)
						if err != nil {
							log.Fatal(err)
							return
						}
						println(string(raw))
						return
					},
				},
				{
					Name:  "invitation",
					Usage: "invitation new user",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:   "api-key, a",
							Value:  "api-key",
							Usage:  "api key for SORACOM",
							EnvVar: "SORACOM_API_KEY",
						},
						cli.StringFlag{
							Name:   "token, t",
							Value:  "token",
							Usage:  "token for SORACOM",
							EnvVar: "SORACOM_TOKEN",
						},
						cli.StringFlag{
							Name:  "email, e",
							Value: "email",
							Usage: "email for invitaiton",
						},
						cli.StringFlag{
							Name:  "password, p",
							Value: "password",
							Usage: "password for invitation user",
						},
					},
					Action: func(c *cli.Context) {
						_, raw, err := operator.ExecOperators(c)
						if err != nil {
							log.Fatal(err)
							return
						}
						println(string(raw))
						return
					},
				},
				{
					Name:  "verify",
					Usage: "verify of invitation user",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:   "api-key, a",
							Value:  "api-key",
							Usage:  "api key for SORACOM",
							EnvVar: "SORACOM_API_KEY",
						},
						cli.StringFlag{
							Name:   "token, t",
							Value:  "token",
							Usage:  "token for SORACOM",
							EnvVar: "SORACOM_TOKEN",
						},
						cli.StringFlag{
							Name:  "verify-token",
							Value: "verify-token",
							Usage: "verify token for intivation user",
						},
					},
					Action: func(c *cli.Context) {
						_, raw, err := operator.ExecVerify(c)
						if err != nil {
							log.Fatal(err)
							return
						}
						println(string(raw))
						return
					},
				},
				{
					Name:  "info",
					Usage: "get information of operator",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:   "api-key, a",
							Value:  "api-key",
							Usage:  "api key for SORACOM",
							EnvVar: "SORACOM_API_KEY",
						},
						cli.StringFlag{
							Name:   "token, t",
							Value:  "token",
							Usage:  "token for SORACOM",
							EnvVar: "SORACOM_TOKEN",
						},
						cli.StringFlag{
							Name:  "operator-id, i",
							Value: "operator-id",
							Usage: "Set operator id",
						},
					},
					Action: func(c *cli.Context) {
						_, raw, err := operator.ExecOperatorInfo(c)
						if err != nil {
							log.Fatal(err)
							return
						}
						println(string(raw))
						return
					},
				},
			},
		},
	}
}
