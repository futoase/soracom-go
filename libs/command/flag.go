package command

import (
  "github.com/codegangsta/cli"
)

func setFlag() []cli.Flag {
  return []cli.Flag {
    cli.BoolFlag{
      Name: "auth, a",
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
    cli.IntFlag{
      Name: "token-timeout-seconds",
      Usage: "token-timeout-seconds",
    },
  }
}
