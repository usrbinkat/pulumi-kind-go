package resources

import (
	"github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes"
	"github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/apiextensions"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func CreateCertManager(ctx *pulumi.Context, kubeProvider *kubernetes.Provider) (*apiextensions.CustomResource, error) {
	// Deploy the Cert Manager
	manager, err := apiextensions.NewCustomResource(ctx, "cert-manager", &apiextensions.CustomResourceArgs{
		ApiVersion: pulumi.String("cert-manager.io/v1"),
		Kind:       pulumi.String("CertManager"),
		// Add other specs similar to your TypeScript code
	}, pulumi.Provider(kubeProvider))

	if err != nil {
		return nil, err
	}

	// Create a cluster issuer that uses self-signed certificates
	rootIssuer, err := apiextensions.NewCustomResource(ctx, "issuerRoot", &apiextensions.CustomResourceArgs{
		ApiVersion: pulumi.String("cert-manager.io/v1"),
		Kind:       pulumi.String("ClusterIssuer"),
		// Add other specs similar to your TypeScript code
	}, pulumi.Provider(kubeProvider), pulumi.DependsOn([]pulumi.Resource{manager}))

	if err != nil {
		return nil, err
	}

	// Create Self Signed ClusterIssuer
	// TODO: Expand this to support using a real CA
	selfSignIssuer, err := apiextensions.NewCustomResource(ctx, "selfSignIssuer", &apiextensions.CustomResourceArgs{
		ApiVersion: pulumi.String("cert-manager.io/v1"),
		Kind:       pulumi.String("ClusterIssuer"),
		// Add other specs similar to your TypeScript code
	}, pulumi.Provider(kubeProvider), pulumi.DependsOn([]pulumi.Resource{rootIssuer}))

	if err != nil {
		return nil, err
	}

	return selfSignIssuer, nil
}
