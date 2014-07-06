package main

import (
	"os"

	"github.com/codegangsta/cli"
)

var Version string = "HEAD"

func main() {
	newApp().Run(os.Args)
}

func newApp() *cli.App {
	app := cli.NewApp()
	app.Name = "ddldoc"
	app.Usage = "Generates documentation files from DDL"
	app.Version = Version
	app.Author = "Yuya Takeyama"
	app.Email = "sign.of.the.wolf.pentagram@gmail.com"
	app.Commands = Commands
	return app
}
