package command

import (
	"github.com/urfave/cli/v2"
)

var WorkflowCommand *cli.Command = &cli.Command{
	Name:        "workflow",
	Description: "Managing workflows",
	Usage:       "Managing workflows",
	Subcommands: []*cli.Command{
		WorkflowListCommand,
		WorkflowExecCommand,
		WorkflowInfoCommand,
	},
}
