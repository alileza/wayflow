package command

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/rodaine/table"
	"github.com/urfave/cli/v2"

	"github.com/alileza/wayflow/workflow"
)

var WorkflowInfoInputs struct {
	ID string
}

var WorkflowInfoCommand *cli.Command = &cli.Command{
	Name:        "info",
	Description: "Get workflow information detail",
	Usage:       "Get workflow information detail",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:        "id",
			Destination: &WorkflowInfoInputs.ID,
			Required:    true,
		},
	},
	Action: func(c *cli.Context) error {
		wm, err := workflow.NewWorkflowManager("./storage")
		if err != nil {
			return err
		}

		w, err := wm.GetWorkflow(workflow.GetWorkflowOptions{ID: WorkflowInfoInputs.ID})
		if err != nil {
			return err
		}

		fmt.Printf("ID: %s\n", w.ID)
		fmt.Printf("Name: %s\n", w.Name)
		fmt.Printf("Description: %s\n", w.Description)
		fmt.Printf("Version: %s\n", w.Version)
		headerFmt := color.New(color.FgHiMagenta, color.Underline).SprintfFunc()
		columnFmt := color.New(color.Bold).SprintfFunc()
		tbl := table.New("Name", "Mapping", "Dependencies")
		tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
		for _, t := range w.Tasks {
			tbl.AddRow(t.Name, t.Mappings, t.Dependencies)
		}
		tbl.AddRow(w.ID, w.Name, w.Version, w.Description)

		tbl.Print()
		return nil
	},
}
