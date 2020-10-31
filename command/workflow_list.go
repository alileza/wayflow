package command

import (
	"github.com/fatih/color"
	"github.com/rodaine/table"
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

		headerFmt := color.New(color.FgHiMagenta, color.Underline).SprintfFunc()
		columnFmt := color.New(color.Bold).SprintfFunc()
		tbl := table.New("ID", "Name", "Version", "Description")
		tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
		for _, w := range wm.Workflows {
			tbl.AddRow(w.ID, w.Name, w.Version, w.Description)
		}

		tbl.Print()
		return nil
	},
}
