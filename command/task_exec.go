package command

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli/v2"

	"github.com/alileza/wayflow/workflow"
)

var TaskExecInputs struct {
	ID        string
	Arguments cli.StringSlice
}

var TaskExecCommand *cli.Command = &cli.Command{
	Name:        "exec",
	Description: "Execute given task",
	Usage:       "Execute given task",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:        "id",
			Destination: &TaskExecInputs.ID,
			Required:    true,
		},
		&cli.StringSliceFlag{
			Name:        "arg",
			Destination: &TaskExecInputs.Arguments,
		},
	},
	Action: func(c *cli.Context) error {
		tm, err := workflow.NewTaskManager("./storage")
		if err != nil {
			return err
		}

		t, err := tm.GetTask(workflow.GetTaskOptions{ID: TaskExecInputs.ID})
		if err != nil {
			return err
		}

		arguments, err := parseArguments(TaskExecInputs.Arguments.Value(), t.GetExpectedInputs())
		if err != nil {
			return err
		}

		resp, err := tm.Exec(c.Context, t, arguments)
		if err != nil {
			return err
		}

		out, err := json.MarshalIndent(resp, "", "    ")
		if err != nil {
			return err
		}

		fmt.Fprintf(os.Stdout, "%s\n", out)

		return nil
	},
}

func parseArguments(args []string, expectedInputs map[string]struct{}) (workflow.TaskArguments, error) {
	result := make(workflow.TaskArguments)
	for _, arg := range args {
		a := strings.Split(arg, "=")
		if len(a) != 2 {
			return nil, fmt.Errorf("Invalid args format of: %v", arg)
		}
		result[a[0]] = a[1]
	}

	for key := range expectedInputs {
		if _, ok := result[key]; !ok {
			return nil, fmt.Errorf("Missing argument: %s", key)
		}
	}

	// for key := range result {
	// 	if _, ok := expectedInputs[key]; !ok {
	// 		return nil, fmt.Errorf("Unexpected argument: %s", key)
	// 	}
	// }
	return result, nil
}
