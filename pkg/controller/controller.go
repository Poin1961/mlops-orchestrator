package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Poin1961/mlops-orchestrator/pkg/api"
	"github.com/Poin1961/mlops-orchestrator/pkg/engine"
)

type WorkflowController struct {
	engine *engine.WorkflowEngine
}

func NewWorkflowController(e *engine.WorkflowEngine) *WorkflowController {
	return &WorkflowController{
		engine: e,
	}
}

func (wc *WorkflowController) CreateWorkflowHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var req api.CreateWorkflowRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	workflow := wc.engine.CreateWorkflow(req.Name, req.Description, req.Tasks)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(workflow)
}

func (wc *WorkflowController) StartWorkflowHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	workflowID := r.URL.Query().Get("id")
	if workflowID == "" {
		http.Error(w, "Workflow ID is required", http.StatusBadRequest)
		return
	}

	err := wc.engine.StartWorkflow(workflowID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Workflow %s started successfully", workflowID)
}

func (wc *WorkflowController) GetWorkflowStatusHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}

	workflowID := r.URL.Query().Get("id")
	if workflowID == "" {
		http.Error(w, "Workflow ID is required", http.StatusBadRequest)
		return
	}

	workflow, err := wc.engine.GetWorkflowStatus(workflowID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(workflow)
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "MLOps Orchestrator is healthy!")
}
