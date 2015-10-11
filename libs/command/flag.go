package command

import (
  "github.com/codegangsta/cli"
)

func setFlag() []cli.Flag {
  return []cli.Flag {
    cli.BoolFlag{
      Name: "auth, a",
    },
    cli.BoolFlag{
      Name: "password_reset_token",
    },
    cli.BoolFlag{
      Name: "issue",
    },
    cli.BoolFlag{
      Name: "verify",
    },
    cli.StringFlag{
      Name: "api_key",
      Value: "api_key",
      EnvVar: "X-SORACOM-API-KEY",
    },
    cli.StringFlag{
      Name: "token",
      Value: "token",
      EnvVar: "X-SORACOM-TOKEN",
    },
    cli.StringFlag{
      Name: "email, e",
      Value: "email",
      EnvVar: "SORACOM_EMAIL",
    },
    cli.StringFlag{
      Name: "password, p",
      Value: "password",
      EnvVar: "SORACOM_PASSWORD",
    },
    cli.StringFlag{
      Name: "onetime_token",
      Value: "onetime_token",
    },
    cli.IntFlag{
      Name: "token-timeout-seconds",
      Usage: "token-timeout-seconds",
    },
  }
}
