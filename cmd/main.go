package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Poin1961/mlops-orchestrator/pkg/api"
	"github.com/Poin1961/mlops-orchestrator/pkg/controller"
	"github.com/Poin1961/mlops-orchestrator/pkg/engine"
)

func main() {
	fmt.Println("Starting MLOps Workflow Orchestrator...")

	// Initialize the workflow engine
	wfEngine := engine.NewWorkflowEngine()

	// Initialize the controller with the engine
	wfController := controller.NewWorkflowController(wfEngine)

	// Setup API routes
	http.HandleFunc("/health", api.HealthCheckHandler)
	http.HandleFunc("/workflow/create", wfController.CreateWorkflowHandler)
	http.HandleFunc("/workflow/start", wfController.StartWorkflowHandler)
	http.HandleFunc("/workflow/status", wfController.GetWorkflowStatusHandler)

	// Start HTTP server in a goroutine
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server := &http.Server{
		Addr: ":" + port,
	}

	go func() {
		fmt.Printf("Server listening on :%s\n", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on %s: %v\n", port, err)
		}
	}()

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop // Wait for interrupt signal

	fmt.Println("Shutting down server...")
	if err := server.Shutdown(nil); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
	fmt.Println("Server gracefully stopped.")
}
