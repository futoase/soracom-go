package command

import (
	"os"
)

func Run() {
	app := InitializeApp()
	SetAction(app)
	app.Run(os.Args)
}
