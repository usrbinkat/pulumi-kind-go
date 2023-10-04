package resources

import (
	"github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func CreateNamespace(ctx *pulumi.Context, name string, kubeProvider *kubernetes.Provider) (*corev1.Namespace, error) {
	ns, err := corev1.NewNamespace(ctx, name, &corev1.NamespaceArgs{
		Metadata: metav1.ObjectMetaArgs{
			Name: pulumi.StringPtr(name),
		},
	}, pulumi.Provider(kubeProvider))
	if err != nil {
		return nil, err
	}
	return ns, nil
}
