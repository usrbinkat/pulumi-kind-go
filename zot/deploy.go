// Zot/deploy.go
package zot

import (
	v1apps "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/apps/v1"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func DeployZotContainer(ctx *pulumi.Context, opts ...pulumi.ResourceOption) (*v1apps.Deployment, error) {
	ctx.Log.Info("Deploying Zot container...", nil)

	deploymentRes, err := v1apps.NewDeployment(ctx, "zot-deployment", &v1apps.DeploymentArgs{
		Spec: v1apps.DeploymentSpecArgs{
			Selector: metav1.LabelSelectorArgs{
				MatchLabels: pulumi.StringMap{
					"app": pulumi.String("zot"),
				},
			},
			Template: corev1.PodTemplateSpecArgs{
				Metadata: metav1.ObjectMetaArgs{
					Labels: pulumi.StringMap{
						"app": pulumi.String("zot"),
					},
				},
				Spec: corev1.PodSpecArgs{
					Containers: corev1.ContainerArray{
						corev1.ContainerArgs{
							Name:  pulumi.String("zot"),
							Image: pulumi.String("ghcr.io/project-zot/zot-linux-amd64:latest"),
						},
					},
				},
			},
		},
	}, opts...)

	if err != nil {
		return nil, err
	}

	ctx.Log.Info("Zot container deployment succeeded", nil)
	return deploymentRes, nil
}
