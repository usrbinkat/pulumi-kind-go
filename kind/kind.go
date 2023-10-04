// File: kind/kind.go
package kind

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pulumi/pulumi-command/sdk/go/command/local"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
	"github.com/usrbinkat/pulumi-kind-go/helper"
)

// ConfigureAndCreateCluster initializes the Kind cluster configuration and creates the cluster.
func ConfigureAndCreateCluster(ctx *pulumi.Context) (*helper.KindCluster, error) {
	ctx.Log.Info("Starting Kind cluster configuration...", nil)

	// Initialize Pulumi config
	cfg := config.New(ctx, "")

	// Load Kind cluster configuration
	kindArgs, err := helper.LoadConfig(cfg)
	if err != nil {
		return nil, fmt.Errorf("Failed to load Kind cluster configuration: %w", err)
	}

	// Create the Kind cluster
	kindClusterResource, err := CreateKindCluster(ctx, kindArgs)
	if err != nil {
		return nil, fmt.Errorf("Error occurred during Kind cluster creation: %w", err)
	}

	ctx.Log.Info("Kind cluster configuration is ready", nil)
	return kindClusterResource, nil
}

// CreateKindCluster is an entry point function for creating a Kind cluster.
func CreateKindCluster(ctx *pulumi.Context, args *helper.KindClusterArgs) (*helper.KindCluster, error) {
	return NewKindCluster(ctx, args.ClusterName, args)
}

// NewKindCluster creates and manages a Kind cluster.
func NewKindCluster(ctx *pulumi.Context, name string, args *helper.KindClusterArgs) (*helper.KindCluster, error) {
	// Default Docker volumes; can be overridden by config or env vars
	volumeNames := []string{"kind-worker1-containerd", "kind-control1-containerd"}
	if err := manageDockerVolumes(ctx, name, args, volumeNames); err != nil {
		return nil, fmt.Errorf("Failed to manage Docker volumes: %w", err)
	}

	// Register the component
	component := &helper.KindCluster{}
	if err := ctx.RegisterComponentResource("my:kind:KindCluster", name, component); err != nil {
		return nil, fmt.Errorf("Failed to register component: %w", err)
	}

	// Create and delete the Kind cluster
	if err := manageKindCluster(ctx, name, args, component); err != nil {
		return nil, fmt.Errorf("Error in managing Kind cluster: %w", err)
	}

	return component, nil
}

// create and deletes Kind cluster
func manageKindCluster(ctx *pulumi.Context, name string, args *helper.KindClusterArgs, component *helper.KindCluster) error {
	cfg := config.New(ctx, "")
	configFile := cfg.Get("kindConfig")

	// Validate and resolve the config file path
	if configFile == "" {
		configFile = "./kind/config.yaml"
		configFile, _ = filepath.Abs(configFile) // Resolve to absolute path
		if _, err := os.Stat(configFile); os.IsNotExist(err) {
			ctx.Log.Warn("No Kind configuration file found; proceeding without --config flag", nil)
			configFile = ""
		}
	} else {
		configFile, _ = filepath.Abs(configFile) // Resolve to absolute path
		if _, err := os.Stat(configFile); os.IsNotExist(err) {
			return fmt.Errorf("Config file specified by pulumi kindConfig does not exist: %s", configFile)
		}
	}

	// Build the Kind commands
	createClusterCmd := buildKindCmd(args, "create", configFile)
	deleteClusterCmd := buildKindCmd(args, "delete", "")

	ctx.Log.Debug(fmt.Sprintf("Executing command: %s", createClusterCmd), nil) // Verbose logging

	// Create the Kind cluster
	createCluster, err := local.NewCommand(ctx, name+"-createCluster", &local.CommandArgs{
		Create: pulumi.StringPtr(createClusterCmd),
		Dir:    pulumi.StringPtr(args.WorkingDir),
	}, pulumi.Parent(component))
	if err != nil {
		return fmt.Errorf("Failed to create Kind cluster: %w", err)
	}

	// Delete the Kind cluster (for cleanup)
	_, err = local.NewCommand(ctx, name+"-deleteCluster", &local.CommandArgs{
		Delete: pulumi.StringPtr(deleteClusterCmd),
		Dir:    pulumi.StringPtr(args.WorkingDir),
	}, pulumi.Parent(component))
	if err != nil {
		return fmt.Errorf("Failed to delete Kind cluster: %w", err)
	}

	// Populate component outputs
	component.ClusterName = pulumi.ToOutput(pulumi.String(args.ClusterName)).(pulumi.StringOutput)
	component.CreateStdout = createCluster.Stdout
	component.DeleteStdout = createCluster.Stdout

	// Register resource outputs
	if err := ctx.RegisterResourceOutputs(component, pulumi.Map{
		"clusterName":  component.ClusterName,
		"createStdout": component.CreateStdout,
		"deleteStdout": component.DeleteStdout,
	}); err != nil {
		return fmt.Errorf("Failed to register resource outputs: %w", err)
	}

	return nil
}

// buildKindCmd constructs the Kind command string based on the provided arguments and config file path.
func buildKindCmd(args *helper.KindClusterArgs, action string, configFile string) string {
	cmd := fmt.Sprintf("kind %s cluster --name %s", action, args.ClusterName)
	if configFile != "" {
		cmd += fmt.Sprintf(" --config %s", configFile)
	}
	return cmd
}
