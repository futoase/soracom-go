package command

import (
  "github.com/codegangsta/cli"
)

func InitializeApp() *cli.App {
  app := cli.NewApp()
  app.Name = "soracom"
  app.Usage = "soracom cli tool for go"
  app.Flags = setFlag()

  return app
}
