// File: cert-manager/init/provider.go
// Use: code to initialize the Kubernetes provider.
package utils

import (
	"github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func InitKubeProvider(ctx *pulumi.Context) (*kubernetes.Provider, error) {
	// Read kubeconfig and context from environment or other configuration
	kubeConfig := GetConfigValue(ctx, "kubeconfig")
	kubeContext := GetConfigValue(ctx, "context")

	// Initialize Kubernetes Provider
	kubeProvider, err := kubernetes.NewProvider(ctx, "kubeconfig", &kubernetes.ProviderArgs{
		Kubeconfig: pulumi.String(kubeConfig),
		Context:    pulumi.String(kubeContext),
	})
	if err != nil {
		return nil, err
	}

	return kubeProvider, nil
}
