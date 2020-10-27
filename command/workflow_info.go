package command

import (
	"fmt"

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

		fmt.Println("ID:", w.ID)
		fmt.Println("Name:", w.Name)
		fmt.Println("Description:", w.Description)

		// g := workflow.NewGraph(w)
		// fmt.Printf("%s\n", g.Dot(&dag.DotOpts{}))
		// fmt.Printf("Graph: %v\n", w.Tasks)

		return nil
	},
}
