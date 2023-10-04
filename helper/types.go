// File: helper/types.go
package helper

import "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

// KindClusterArgs holds the arguments for creating a new Kind cluster.
// ClusterName is the name of the Kind cluster.
// WorkingDir specifies the working directory for the Kind config.yaml file.
// Purge indicates whether to remove pre-existing resources related to this Kind cluster.
type KindClusterArgs struct {
	ClusterName string
	WorkingDir  string
	Purge       bool
}

// KindCluster represents a Pulumi component for managing a Kind cluster.
// ClusterName is the name of the Kind cluster.
// CreateStdout captures the standard output of the Kind cluster creation process.
// DeleteStdout captures the standard output of the Kind cluster deletion process.
type KindCluster struct {
	pulumi.ResourceState

	ClusterName  pulumi.StringOutput
	CreateStdout pulumi.StringOutput
	DeleteStdout pulumi.StringOutput
}
