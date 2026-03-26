package controller

import (
	"fmt"
	"io/ioutil"
	"log"
	"mlops-orchestrator/pkg/api"
	"mlops-orchestrator/pkg/engine"

	"gopkg.in/yaml.v2"
)

type Controller struct {
	workflow api.Workflow
	engine   *engine.Engine
}

func NewController(workflowFile string) (*Controller, error) {
	data, err := ioutil.ReadFile(workflowFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read workflow file: %w", err)
	}

	var workflow api.Workflow
	if err := yaml.Unmarshal(data, &workflow); err != nil {
		return nil, fmt.Errorf("failed to unmarshal workflow YAML: %w", err)
	}

	eng := engine.NewEngine()

	return &Controller{
		workflow: workflow,
		engine:   eng,
	}, nil
}

func (c *Controller) Run() error {
	log.Printf("Executing workflow: %s\n", c.workflow.Metadata.Name)

	for _, task := range c.workflow.Spec.Tasks {
		log.Printf("Running task: %s\n", task.Name)
		if err := c.engine.RunTask(task); err != nil {
			return fmt.Errorf("task %s failed: %w", task.Name, err)
		}
		log.Printf("Task %s completed.\n", task.Name)
	}

	return nil
}
