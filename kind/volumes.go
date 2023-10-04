// File: kind/volumes.go
package kind

import (
	"fmt"
	"strings"

	"github.com/pulumi/pulumi-command/sdk/go/command/local"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/usrbinkat/pulumi-kind-go/helper"
)

func manageDockerVolumes(ctx *pulumi.Context, name string, args *helper.KindClusterArgs, volumeNames []string) error {
	volumeExistsCmd := fmt.Sprintf("docker volume ls --filter name=%s -q", strings.Join(volumeNames, " --filter name="))
	volumeCheck, err := local.NewCommand(ctx, name+"-volumeCheck", &local.CommandArgs{
		Create: pulumi.StringPtr(volumeExistsCmd),
	})
	if err != nil {
		return fmt.Errorf("Failed to check Docker volumes: %w", err)
	}
	output := volumeCheck.Stdout.ToStringOutput().ApplyT(func(s string) []string { return strings.Fields(s) }).(pulumi.StringArrayOutput)

	// Use ApplyT to get the string array from the output and proceed with the logic
	output.ApplyT(func(arr []string) interface{} {
		existingVolumes := arr
		// Delete if purge=true
		if args.Purge {
			deleteCmd := "docker volume rm " + strings.Join(existingVolumes, " ")
			_, err := local.NewCommand(ctx, name+"-deleteVolumes", &local.CommandArgs{
				Create: pulumi.StringPtr(deleteCmd),
			})
			if err != nil {
				return fmt.Errorf("Failed to delete existing Docker volumes: %w", err)
			}
		}

		// Create missing volumes
		for _, volume := range volumeNames {
			if !strings.Contains(strings.Join(existingVolumes, " "), volume) {
				createCmd := "docker volume create " + volume
				_, err := local.NewCommand(ctx, name+"-create-"+volume, &local.CommandArgs{
					Create: pulumi.StringPtr(createCmd),
				})
				if err != nil {
					return fmt.Errorf("Failed to create Docker volume: %w", err)
				}
			}
		}
		return nil // return nil if nothing goes wrong
	})

	return nil
}
