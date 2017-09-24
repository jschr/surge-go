package main

import (
	"github.com/jschr/surge-go"
	"github.com/segmentio/go-prompt"
	"github.com/urfave/cli"
)

// Publish the project to the domain domain
func Publish(c *cli.Context, s surge.Surge) error {
	err := Login(c, s)
	if err != nil {
		return err
	}

	projectPath := c.String("project")
	if projectPath == "" {
		projectPath = prompt.String("project")
	}

	domain := c.String("domain")
	if domain == "" {
		domain = prompt.String("domain")
	}

	err = s.Publish(projectPath, domain)
	if err != nil {
		return cli.NewExitError(err, 1)
	}

	return nil
}
