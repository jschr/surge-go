package main

import (
	"github.com/jschr/surge-go"
	"github.com/segmentio/go-prompt"
	"github.com/urfave/cli"
)

// Login runs the login cli command
func Login(c *cli.Context, s surge.Surge) error {
	username := c.String("username")
	if username == "" {
		username = prompt.String("username")
	}

	password := c.String("password")
	if password == "" {
		password = prompt.PasswordMasked("password")
	}

	token, err := s.Login(username, password)
	if err != nil {
		return cli.NewExitError(err, 1)
	}

	s.SetToken(token)

	return nil
}
