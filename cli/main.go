package main

import (
	"os"

	"github.com/jschr/surge-go"
	"github.com/urfave/cli"
)

// Flags contains all options for the cli
type Flags struct {
	Username cli.StringFlag
	Password cli.StringFlag
	Project  cli.StringFlag
	Domain   cli.StringFlag
	// Add      cli.StringFlag
	// Remove   cli.StringFlag
}

func main() {
	s := surge.NewSurge("")
	app := cli.NewApp()

	app.Name = "surge"
	app.Usage = "Unofficial CLI for the surge.sh CDN"

	flags := &Flags{
		cli.StringFlag{
			Name:   "username, u",
			EnvVar: "SURGE_USERNAME",
			Usage:  "Your surge `username`",
		},
		cli.StringFlag{
			Name:   "password, pw",
			EnvVar: "SURGE_PASSWORD",
			Usage:  "Your surge `password`",
		},
		cli.StringFlag{
			Name:   "project, p",
			EnvVar: "SURGE_PROJECT",
			Usage:  "Path to projects asset directory",
			Value:  "./",
		},
		cli.StringFlag{
			Name:   "domain, d",
			EnvVar: "SURGE_DOMAIN",
			Usage:  "Your surge `password`",
		},
	}
	app.Flags = []cli.Flag{flags.Username, flags.Password, flags.Project, flags.Domain}

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
		{
			Name:  "teardown",
			Usage: "tear down a published project",
			Flags: []cli.Flag{flags.Username, flags.Password},
			Action: func(c *cli.Context) error {
				return Teardown(c, s)
			},
		},
	}

	app.Action = func(c *cli.Context) error {
		return Publish(c, s)
	}

	app.Run(os.Args)
}
