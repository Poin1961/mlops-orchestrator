package engine

import (
	"fmt"
	"log"
	"mlops-orchestrator/pkg/api"
	"os/exec"
)

type Engine struct {
	// Add any engine-specific state here
}

func NewEngine() *Engine {
	return &Engine{}
}

func (e *Engine) RunTask(task api.Task) error {
	log.Printf("Executing command for task %s: %v\n", task.Name, task.Command)

	cmd := exec.Command(task.Command[0], task.Command[1:]...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("command execution failed for task %s: %w\nOutput: %s", task.Name, err, string(output))
	}

	log.Printf("Task %s output:\n%s", task.Name, string(output))

	// Simulate artifact handling (e.g., copying data from/to S3)
	for _, input := range task.Inputs {
		log.Printf("Simulating input artifact: %s from %s\n", input.Name, input.From)
	}
	for _, output := range task.Outputs {
		log.Printf("Simulating output artifact: %s to %s\n", output.Name, output.To)
	}

	return nil
}
