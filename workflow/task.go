package workflow

import (
	"fmt"
	"io/ioutil"
	"os"

	docker "github.com/docker/docker/client"
	"gopkg.in/yaml.v2"
)

type Task struct {
	Version     string   `yaml:"version" json:"version"`
	ID          string   `yaml:"id"  json:"id"`
	Name        string   `yaml:"name" json:"name"`
	Description string   `yaml:"description" json:"description"`
	Inputs      []string `yaml:"inputs" json:"inputs"`
	Outputs     []string `yaml:"outputs" json:"outputs"`
	Run         struct {
		Provider string `yaml:"provider" json:"provider"`
		Handler  string `yaml:"handler" json:"handler"`
	} `yaml:"run" json:"run"`
}

func (t *Task) GetExpectedInputs() map[string]struct{} {
	expectedResult := make(map[string]struct{})
	for _, input := range t.Inputs {
		expectedResult[input] = struct{}{}
	}
	return expectedResult
}

type TaskArguments map[string]string

type TaskManager struct {
	Tasks        []*Task
	DockerClient *docker.Client
}

func NewTaskManager(storagePath string) (*TaskManager, error) {
	var tm TaskManager

	files, err := ioutil.ReadDir(storagePath + "/tasks")
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		t, err := readTask(storagePath + "/tasks/" + file.Name())
		if err != nil {
			return nil, err
		}
		tm.Tasks = append(tm.Tasks, t)
	}

	tm.DockerClient, err = docker.NewEnvClient()
	if err != nil {
		return nil, err
	}

	return &tm, nil
}

type GetTaskOptions struct {
	ID string
}

func (tm *TaskManager) GetTask(opts GetTaskOptions) (*Task, error) {
	for _, t := range tm.Tasks {
		if t.ID == opts.ID {
			return t, nil
		}
	}
	return nil, fmt.Errorf("task %w: %+v", ErrNotFound, opts)
}

func readTask(path string) (*Task, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var t Task
	if err := yaml.NewDecoder(f).Decode(&t); err != nil {
		return nil, err
	}
	return &t, nil
}
