package main

import (
	"os"

	"github.com/jschr/surge-go"
	"github.com/urfave/cli"
)

type Flags struct {
	Username cli.StringFlag
	Password cli.StringFlag
}

func main() {
	app := cli.NewApp()
	s := surge.NewSurge("")

	flags := &Flags{
		cli.StringFlag{
			Name:   "username, u",
			EnvVar: "SURGE_USERNAME",
			Usage:  "Your surge `username`",
		},
		cli.StringFlag{
			Name:   "password, p",
			EnvVar: "SURGE_PASSWORD",
			Usage:  "Your surge `password`",
		},
	}

	app.Name = "surge"
	app.Usage = "CLI for the surge.sh CDN"
	app.Flags = []cli.Flag{flags.Username, flags.Password}

	app.Commands = []cli.Command{
		{
			Name:  "login",
			Usage: "only performs authentication step",
			Flags: []cli.Flag{flags.Username, flags.Password},
			Action: func(c *cli.Context) error {
				return Login(c, s)
			},
		},
		{
			Name:  "list",
			Usage: "list all domains you have access to",
			Flags: []cli.Flag{flags.Username, flags.Password},
			Action: func(c *cli.Context) error {
				return List(c, s)
			},
		},
	}

	app.Action = func(c *cli.Context) error {
		err := Login(c, s)
		if err != nil {
			return err
		}
		return nil
	}

	app.Run(os.Args)
}
