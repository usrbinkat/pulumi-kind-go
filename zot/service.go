// File: zot/service.go
package zot

import (
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/core/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ExposeZotService(ctx *pulumi.Context, opts ...pulumi.ResourceOption) (*corev1.Service, error) {
	ctx.Log.Info("Publishing Zot service...", nil)

	servicePorts := corev1.ServicePortArray{}
	servicePorts = append(servicePorts, &corev1.ServicePortArgs{
		Port:       pulumi.Int(5000),
		TargetPort: pulumi.Int(5000),
		NodePort:   pulumi.Int(30000),
	})

	serviceRes, err := corev1.NewService(ctx, "zot-service", &corev1.ServiceArgs{
		Spec: corev1.ServiceSpecArgs{
			Selector: pulumi.StringMap{
				"app": pulumi.String("zot"),
			},
			Ports: servicePorts,
			Type:  pulumi.String("NodePort"),
		},
	}, opts...)

	if err != nil {
		return nil, err
	}

	ctx.Log.Info("Zot OCI Registry Service generated", nil)
	return serviceRes, nil
}
