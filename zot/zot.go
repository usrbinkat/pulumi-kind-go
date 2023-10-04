// File: zot/zot.go
package zot

import (
	"fmt"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func DeployZotRegistry(ctx *pulumi.Context, opts ...pulumi.ResourceOption) error {
	ctx.Log.Info("Planning Zot OCI Registry deployment...", nil)

	deploymentRes, err := DeployZotContainer(ctx, opts...)
	if err != nil {
		return fmt.Errorf("Failed during Zot container deployment: %w", err)
	}

	// Make service dependent on the deployment
	serviceResOpts := append(opts, pulumi.DependsOn([]pulumi.Resource{deploymentRes}))
	if _, err := ExposeZotService(ctx, serviceResOpts...); err != nil {
		return fmt.Errorf("Failed to configure Zot service: %w", err)
	}

	ctx.Log.Info("Zot OCI Registry deployment is ready.", nil)
	return nil
}
