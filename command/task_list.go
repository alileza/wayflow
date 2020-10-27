package command

import (
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/alileza/wayflow/workflow"
)

var TaskListCommand *cli.Command = &cli.Command{
	Name:        "list",
	Description: "Get list of tasks",
	Usage:       "Get list of tasks",
	Action: func(c *cli.Context) error {
		tm, err := workflow.NewTaskManager("./storage")
		if err != nil {
			return err
		}
		fmt.Printf("%s %s\t\t\t%s\t%s\n", "NO", "NAME", "VERSION", "DESCRIPTION")
		for d, t := range tm.Tasks {
			fmt.Printf("%d. %s\t%s\t%s\n", d+1, t.Name, t.Version, t.Description)
		}

		return nil
	},
}
