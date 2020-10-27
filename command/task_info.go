package command

import (
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/alileza/wayflow/workflow"
)

var TaskInfoInputs struct {
	ID string
}

var TaskInfoCommand *cli.Command = &cli.Command{
	Name:        "info",
	Description: "Get task information detail",
	Usage:       "Get task information detail",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:        "id",
			Destination: &TaskInfoInputs.ID,
			Required:    true,
		},
	},
	Action: func(c *cli.Context) error {
		tm, err := workflow.NewTaskManager("./storage")
		if err != nil {
			return err
		}

		t, err := tm.GetTask(workflow.GetTaskOptions{ID: TaskInfoInputs.ID})
		if err != nil {
			return err
		}

		fmt.Println("ID:", t.ID)
		fmt.Println("Name:", t.Name)
		fmt.Println("Description:", t.Description)
		fmt.Println("Inputs:", t.Inputs)
		fmt.Println("Outputs:", t.Outputs)
		fmt.Println("Run:")
		fmt.Println("  Provider:", t.Run.Provider)
		fmt.Println("  Handler:", t.Run.Handler)

		return nil
	},
}
