package workflow

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os/exec"
	"time"
)

type TaskExecRequest struct {
	Handler   string        `json:"handler"`
	Arguments TaskArguments `json:"args"`
}

type TaskExecResponse struct {
	Message  string        `json:"message"`
	Inputs   TaskArguments `json:"inputs"`
	Outputs  TaskArguments `json:"outputs"`
	Code     int64         `json:"code"`
	Error    string        `json:"error,omitempty"`
	ExecTime time.Time     `json:"execution_time"`
}

func (tm *TaskManager) Exec(ctx context.Context, t *Task, args TaskArguments) (*TaskExecResponse, error) {
	request := TaskExecRequest{
		Handler:   t.Run.Handler,
		Arguments: args,
	}

	var response TaskExecResponse
	response.ExecTime = time.Now()
	response.Inputs = args
	cmd := exec.Command("docker", "run", "--rm", "-i", t.Run.Provider)
	req, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	cmd.Stdin = bytes.NewReader(req)

	var body bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &body
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("%w: %s", err, stderr.String())
	}

	if err := json.Unmarshal(body.Bytes(), &response); err != nil {
		return nil, fmt.Errorf("unmarshal response error: %w: %s", err, body.String())
	}

	return &response, nil
}
