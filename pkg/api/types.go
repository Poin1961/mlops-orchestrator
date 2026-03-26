package api

type Workflow struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Tasks       []Task `json:"tasks"`
}

type Task struct {
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Type        string            `json:"type"` // e.g., "data_ingestion", "model_training", "model_deployment"
	Status      string            `json:"status"` // e.g., "pending", "running", "completed", "failed"
	Config      map[string]string `json:"config"`
	Dependencies []string         `json:"dependencies"`
}

type CreateWorkflowRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Tasks       []Task `json:"tasks"`
}

type WorkflowStatusResponse struct {
	WorkflowID string `json:"workflow_id"`
	Status     string `json:"status"` // e.g., "running", "completed", "failed"
	Tasks      []Task `json:"tasks"`
}

