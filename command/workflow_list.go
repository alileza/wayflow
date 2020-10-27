package command

import (
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/alileza/wayflow/workflow"
)

var WorkflowListCommand *cli.Command = &cli.Command{
	Name:        "list",
	Description: "Get list of existing workflow",
	Usage:       "List of existing workflows",
	Action: func(c *cli.Context) error {
		wm, err := workflow.NewWorkflowManager("./storage")
		if err != nil {
			return err
		}

		fmt.Printf("%s %s\t\t\t%s\t%s\n", "NO", "NAME", "VERSION", "DESCRIPTION")
		for d, w := range wm.Workflows {
			fmt.Printf("%d. %s\t%s\t%s\n", d+1, w.Name, w.Version, w.Description)
		}

		return nil
	},
}
