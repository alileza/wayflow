package command

import (
	"github.com/fatih/color"
	"github.com/rodaine/table"
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

		headerFmt := color.New(color.FgHiMagenta, color.Underline).SprintfFunc()
		columnFmt := color.New(color.Bold).SprintfFunc()
		tbl := table.New("ID", "Name", "Description")
		tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
		for _, t := range tm.Tasks {
			tbl.AddRow(t.ID, t.Name, t.Description)
		}

		tbl.Print()
		return nil
	},
}
