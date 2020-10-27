package workflow

import (
	"context"
	"fmt"
)

type State string

type WorkflowExecRequest struct {
	WorkflowID string `json:"workflow_id"`
}

type WorkflowExecResponse struct{}

type WorkflowArguments map[string]TaskArguments

func (wm *WorkflowManager) Exec(ctx context.Context, w *Workflow, args WorkflowArguments) (*WorkflowExecResponse, error) {
	stateMap := make(map[string]WorkflowTask)
	for _, t := range w.Tasks {
		stateMap[t.Name] = t
		task, err := wm.TaskManager.GetTask(GetTaskOptions{ID: t.TaskID})
		if err != nil {
			return nil, fmt.Errorf("Can't find task %+v: %w", t, err)
		}

		// Schedule a job
		response, err := wm.TaskManager.Exec(ctx, task, args[t.Name])
		if err != nil {
			return nil, err
		}

		fmt.Printf("%+v\n", response)
	}
	return &WorkflowExecResponse{}, nil
}

//
// workflow
// task
