package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/alileza/wayflow/command"
)

func main() {
	app := &cli.App{
		Name:  "wayflow",
		Usage: "Workflow Engine",
		Commands: []*cli.Command{
			command.WorkflowCommand,
			command.TaskCommand,
			command.JobCommand,
		},
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stdout, "ERR: %v\n", err)
		os.Exit(1)
	}
}
