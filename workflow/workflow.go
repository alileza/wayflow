package workflow

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

var (
	ErrNotFound = errors.New("Not found")
)

// Workflow defines structure o
type Workflow struct {
	Version     string         `yaml:"version"`
	ID          string         `yaml:"id"`
	Name        string         `yaml:"name"`
	Description string         `yaml:"description"`
	Tasks       []WorkflowTask `yaml:"tasks"`
}

type WorkflowTask struct {
	Name         string                    `yaml:"name"`
	TaskID       string                    `yaml:"task_id"`
	State        State                     `yaml:"state"`
	Dependencies []string                  `yaml:"dependencies"`
	Mappings     []WorkflowArgumentMapping `yaml:"mappings"`
}

type ArgumentKey string

func (a ArgumentKey) Key() string {
	l := strings.Split(string(a), ":")
	return l[len(l)-1]
}

type WorkflowArgumentMapping struct {
	From ArgumentKey `yaml:"from"`
	To   ArgumentKey `yaml:"to"`
}

type WorkflowManager struct {
	*TaskManager
	Workflows []*Workflow
}

func NewWorkflowManager(storagePath string) (*WorkflowManager, error) {
	var wm WorkflowManager

	files, err := ioutil.ReadDir(storagePath + "/workflows")
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		w, err := readWorkflow(storagePath + "/workflows/" + file.Name())
		if err != nil {
			return nil, err
		}
		wm.Workflows = append(wm.Workflows, w)
	}

	wm.TaskManager, err = NewTaskManager(storagePath)
	if err != nil {
		return nil, err
	}

	return &wm, nil
}

type GetWorkflowOptions struct {
	ID string
}

func (wm *WorkflowManager) GetWorkflow(opts GetWorkflowOptions) (*Workflow, error) {
	for _, w := range wm.Workflows {
		if w.ID == opts.ID {
			return w, nil
		}
	}
	return nil, fmt.Errorf("workflow %w: %+v", ErrNotFound, opts)
}

func readWorkflow(path string) (*Workflow, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var t Workflow
	if err := yaml.NewDecoder(f).Decode(&t); err != nil {
		return nil, err
	}
	return &t, nil
}
