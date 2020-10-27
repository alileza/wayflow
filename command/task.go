package command

import (
	"github.com/urfave/cli/v2"
)

var TaskCommand *cli.Command = &cli.Command{
	Name:        "task",
	Description: "Managing tasks",
	Usage:       "Managing tasks",
	Subcommands: []*cli.Command{
		TaskListCommand,
		TaskInfoCommand,
		TaskExecCommand,
	},
}
