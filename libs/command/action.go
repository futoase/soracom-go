package command

import (
  "log"
  "github.com/codegangsta/cli"
  auth "github.com/futoase/soracom/libs/auth"
)

func SetAction(a *cli.App) {
  a.Action = func(c *cli.Context) {
    switch {
    case c.IsSet("auth"):
      _, raw, err := auth.Exec(c)
      if err != nil {
        log.Fatal(err)
      }
      println(string(raw))
      return
    case c.IsSet("operator"):
      return
    case c.IsSet("subscriber"):
      return
    case c.IsSet("stats"):
      return
    case c.IsSet("group"):
      return
    case c.IsSet("event_handler"):
      return
    }
  }
}
