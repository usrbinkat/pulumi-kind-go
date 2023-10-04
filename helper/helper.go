// File: helper/helper.go
package helper

import (
	"fmt"
	"os/exec"
	"regexp"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

// LoadConfig loads the Kind cluster configuration.
func LoadConfig(cfg *config.Config) (*KindClusterArgs, error) {
	purge := cfg.GetBool("purge")
	cfgKindCluster := cfg.Get("kindCluster")
	if cfgKindCluster == "" {
		cfgKindCluster = "default-cluster"
	}

	// Validate ClusterName
	if err := validateClusterName(cfgKindCluster); err != nil {
		return nil, fmt.Errorf("ClusterName validation failed: %w", err)
	}

	return &KindClusterArgs{
		ClusterName: cfgKindCluster,
		WorkingDir:  "./kind",
		Purge:       purge,
	}, nil
}

// CheckDependencies checks if a specific dependency is installed.
func CheckDependencies(ctx *pulumi.Context, dependency string) error {
	if _, err := exec.LookPath(dependency); err != nil {
		ctx.Log.Error(fmt.Sprintf("Dependency '%s' is missing", dependency), nil)
		return fmt.Errorf("Dependency '%s' is not installed: %w", dependency, err)
	}
	return nil
}

// validateClusterName validates the Kind cluster name based on certain criteria.
// It uses regex to ensure that the cluster name contains only alphanumeric characters and hyphens.
func validateClusterName(clusterName string) error {
	regex := regexp.MustCompile(`^[a-zA-Z0-9-]+$`)
	if !regex.MatchString(clusterName) {
		return fmt.Errorf("Cluster name contains invalid characters; it should only contain alphanumeric characters and hyphens")
	}
	return nil
}
