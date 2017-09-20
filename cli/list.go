package main

import (
	"fmt"

	"github.com/jschr/surge-go"
	"github.com/urfave/cli"
)

// List all domains
func List(c *cli.Context, s surge.Surge) error {
	err := Login(c, s)
	if err != nil {
		return err
	}

	domains, err := s.List()
	if err != nil {
		return cli.NewExitError(err, 1)
	}

	for _, domain := range domains {
		fmt.Println(domain)
	}

	return nil
}
