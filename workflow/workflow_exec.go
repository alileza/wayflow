package workflow

import (
	"context"
	"fmt"
)

type WorkflowExecRequest struct {
	WorkflowID string `json:"workflow_id"`
}

type WorkflowExecResponse struct {
	TaskExecResponses map[string]*TaskExecResponse `yaml:"task_exec_responses"`
}

type WorkflowArguments map[string]TaskArguments

func (wm *WorkflowManager) ExecWorkflow(ctx context.Context, w *Workflow, args WorkflowArguments) (*WorkflowExecResponse, error) {
	response := &WorkflowExecResponse{
		TaskExecResponses: make(map[string]*TaskExecResponse),
	}

	for _, t := range w.Tasks {
		task, err := wm.TaskManager.GetTask(GetTaskOptions{ID: t.TaskID})
		if err != nil {
			return nil, fmt.Errorf("Can't find task %+v: %w", t, err)
		}

		for _, d := range t.Dependencies {
			resp, ok := response.TaskExecResponses[d]
			if !ok {
				return nil, fmt.Errorf("Failed to get dependency response for task %+v: %s", t, d)
			}
			for _, m := range t.Mappings {
				if m.From.TaskName() != d {
					continue
				}
				v, ok := resp.Outputs[m.From.Key()]
				if !ok {
					return nil, fmt.Errorf("Failed to get key output of %s from task output %s => outputs=%+v", m.From.Key(), d, resp.Outputs)
				}
				args[t.Name][m.To.Key()] = v
			}
		}

		response.TaskExecResponses[t.Name], err = wm.TaskManager.Exec(ctx, task, args[t.Name])
		if err != nil {
			return nil, fmt.Errorf("Failed to execute task %s: %w", task.ID, err)
		}
	}

	return response, nil
}
