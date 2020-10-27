package command

import (
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/alileza/wayflow/workflow"
)

var WorkflowExecInputs struct {
	ID        string
	Arguments cli.StringSlice
}

var WorkflowExecCommand *cli.Command = &cli.Command{
	Name:        "exec",
	Description: "Execute workflow",
	Usage:       "Execute workflow",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:        "id",
			Destination: &WorkflowExecInputs.ID,
			Required:    true,
		},
		&cli.StringSliceFlag{
			Name:        "arg",
			Destination: &WorkflowExecInputs.Arguments,
		},
	},
	Action: func(c *cli.Context) error {
		wm, err := workflow.NewWorkflowManager("./storage")
		if err != nil {
			return err
		}

		w, err := wm.GetWorkflow(workflow.GetWorkflowOptions{ID: WorkflowExecInputs.ID})
		if err != nil {
			return err
		}

		workflowArgument := make(workflow.WorkflowArguments)
		for _, wt := range w.Tasks {
			task, err := wm.TaskManager.GetTask(workflow.GetTaskOptions{ID: wt.TaskID})
			if err != nil {
				return err
			}

			arguments, err := parseArguments(WorkflowExecInputs.Arguments.Value(), task.GetExpectedInputs())
			if err != nil {
				return err
			}
			workflowArgument[wt.Name] = arguments
		}

		r, err := wm.Exec(c.Context, w, workflowArgument)
		if err != nil {
			return err
		}
		fmt.Printf("%+v", r)
		return nil
	},
}
