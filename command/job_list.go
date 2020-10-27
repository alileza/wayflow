package command

import (
	"github.com/urfave/cli/v2"
)

var JobListCommand *cli.Command = &cli.Command{
	Name:        "list",
	Description: "Get list of jobs",
	Usage:       "Get list of jobs",
	Action: func(c *cli.Context) error {
		return nil
	},
}
