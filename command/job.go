package command

import (
	"github.com/urfave/cli/v2"
)

var JobCommand *cli.Command = &cli.Command{
	Name:        "job",
	Description: "Managing job",
	Usage:       "Managing job",
	Subcommands: []*cli.Command{
		JobListCommand,
	},
}
