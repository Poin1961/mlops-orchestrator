package engine

import (
	"fmt"
	"log"
	"sync"

	"github.com/Poin1961/mlops-orchestrator/pkg/api"
	"github.com/google/uuid"
)

type WorkflowEngine struct {
	wfStore map[string]*api.Workflow
	mu      sync.RWMutex
}

func NewWorkflowEngine() *WorkflowEngine {
	return &WorkflowEngine{
		wfStore: make(map[string]*api.Workflow),
	}
}

func (we *WorkflowEngine) CreateWorkflow(name, description string, tasks []api.Task) *api.Workflow {
	we.mu.Lock()
	defer we.mu.Unlock()

	id := uuid.New().String()
	workflow := &api.Workflow{
		ID:          id,
		Name:        name,
		Description: description,
		Tasks:       tasks,
	}

	for i := range workflow.Tasks {
		workflow.Tasks[i].ID = uuid.New().String()
		workflow.Tasks[i].Status = "pending"
	}

	we.wfStore[id] = workflow
	log.Printf("Workflow %s created: %s\n", id, name)
	return workflow
}

func (we *WorkflowEngine) StartWorkflow(workflowID string) error {
	we.mu.Lock()
	defer we.mu.Unlock()

	workflow, ok := we.wfStore[workflowID]
	if !ok {
		return fmt.Errorf("workflow with ID %s not found", workflowID)
	}

	log.Printf("Starting workflow %s\n", workflow.Name)

	// Simulate asynchronous task execution
	go we.executeWorkflow(workflow)

	return nil
}

func (we *WorkflowEngine) executeWorkflow(workflow *api.Workflow) {
	// Simple dependency resolution and execution
	// In a real system, this would be a more sophisticated DAG execution engine

	completedTasks := make(map[string]bool)
	for {
		allTasksCompleted := true
		for i := range workflow.Tasks {
			task := &workflow.Tasks[i]

			if task.Status == "pending" {
				// Check dependencies
				canExecute := true
				for _, depID := range task.Dependencies {
					if !completedTasks[depID] {
						canExecute = false
						break
					}
				}

				if canExecute {
					task.Status = "running"
					log.Printf("Executing task %s (%s) in workflow %s\n", task.Name, task.ID, workflow.Name)
					// Simulate task work
					// time.Sleep(time.Duration(rand.Intn(5)+1) * time.Second)

					// For simplicity, tasks always succeed
					task.Status = "completed"
					completedTasks[task.ID] = true
					log.Printf("Task %s (%s) completed in workflow %s\n", task.Name, task.ID, workflow.Name)
				}
			}

			if task.Status != "completed" && task.Status != "failed" {
				allTasksCompleted = false
			}
		}

		if allTasksCompleted {
			break
		}
		// Small delay to prevent busy-waiting in a real scenario
		// time.Sleep(100 * time.Millisecond)
	}

	log.Printf("Workflow %s (%s) finished.\n", workflow.Name, workflow.ID)
}

func (we *WorkflowEngine) GetWorkflowStatus(workflowID string) (*api.Workflow, error) {
	we.mu.RLock()
	defer we.mu.RUnlock()

	workflow, ok := we.wfStore[workflowID]
	if !ok {
		return nil, fmt.Errorf("workflow with ID %s not found", workflowID)
	}
	return workflow, nil
}
