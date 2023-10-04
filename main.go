// File: main.go
package main

import (
	"fmt"
	"sync"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/usrbinkat/pulumi-kind-go/helper"
	"github.com/usrbinkat/pulumi-kind-go/kind"
	"github.com/usrbinkat/pulumi-kind-go/zot"
)

func main() {
	pulumi.Run(runProgram)
}

func runProgram(ctx *pulumi.Context) error {
	ctx.Log.Info("Starting the initialization process...", nil)

	// Initialize and check for errors with enhanced context
	if err := initialize(ctx); err != nil {
		return fmt.Errorf("Initialization failed during dependency check: %w", err)
	}
	ctx.Log.Info("Initialization and dependency checks are complete", nil)

	// Configure and create the Kind cluster
	kindClusterResource, err := kind.ConfigureAndCreateCluster(ctx)
	if err != nil {
		return fmt.Errorf("Failed during Kind cluster creation: %w", err)
	}
	ctx.Log.Info("Ready to build Kind cluster", nil)

	// Deploy Zot OCI registry
	if err := zot.DeployZotRegistry(ctx, pulumi.DependsOn([]pulumi.Resource{kindClusterResource})); err != nil {
		return fmt.Errorf("Failed to deploy Zot OCI registry: %w", err)
	}
	ctx.Log.Info("Zot OCI registry deployed successfully", nil)

	return nil
}

func initialize(ctx *pulumi.Context) error {
	ctx.Log.Info("Starting dependency checks...", nil)

	// Define dependencies
	dependencies := []string{"kind", "docker"}

	// Initialize a WaitGroup for concurrent execution
	var wg sync.WaitGroup
	errChan := make(chan error, len(dependencies))

	// Perform dependency checks concurrently
	for _, dependency := range dependencies {
		wg.Add(1)
		go func(dep string) {
			defer wg.Done()
			if err := helper.CheckDependencies(ctx, dep); err != nil {
				errChan <- fmt.Errorf("Dependency check failed for %s: %w", dep, err)
			}
		}(dependency)
	}

	// Wait for all goroutines to complete
	wg.Wait()
	close(errChan)

	// Check if any errors occurred during the dependency checks
	for err := range errChan {
		if err != nil {
			return err
		}
	}

	ctx.Log.Info("Successfully detected all local dependencies.", nil)
	return nil
}
