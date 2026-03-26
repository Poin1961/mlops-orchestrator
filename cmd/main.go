package main

import (
	"fmt"
	"log"
	"os"

	"mlops-orchestrator/pkg/controller"
)

func main() {
	if len(os.Args) < 3 || os.Args[1] != "run" {
		fmt.Println("Usage: mlops-orchestrator run <workflow-file.yaml>")
		os.Exit(1)
	}

	workflowFile := os.Args[2]
	fmt.Printf("Starting MLOps Workflow Orchestrator with workflow: %s\n", workflowFile)

	c, err := controller.NewController(workflowFile)
	if err != nil {
		log.Fatalf("Failed to create controller: %v", err)
	}

	if err := c.Run(); err != nil {
		log.Fatalf("Workflow execution failed: %v", err)
	}

	fmt.Println("Workflow completed successfully!")
}
