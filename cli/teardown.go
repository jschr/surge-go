package main

import (
	"github.com/jschr/surge-go"
	"github.com/segmentio/go-prompt"
	"github.com/urfave/cli"
)

// Teardown project at domain
func Teardown(c *cli.Context, s surge.Surge) error {
	err := Login(c, s)
	if err != nil {
		return err
	}

	domain := c.String("domain")
	if domain == "" {
		domain = prompt.String("domain")
	}

	err = s.Teardown(domain)
	if err != nil {
		return err
	}

	return nil
}
