package command

import (
  "log"
  "github.com/codegangsta/cli"
  auth "github.com/futoase/soracom/libs/auth"
)

func SetAction(a *cli.App) {
  a.Action = func(c *cli.Context) {
    if c.IsSet("auth") {
      _, raw, err := auth.Exec(c)
      if err != nil {
        log.Fatal(err)
      }
      println(string(raw))
    }
  }
}
