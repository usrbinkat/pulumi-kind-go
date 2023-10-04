// File: cert-manager/utils/config.go
// Use: code to read configuration values from environment variables or Pulumi configuration.
package utils

import (
	"os"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func GetConfigValue(ctx *pulumi.Context, key string) string {
	// Attempt to read from an environment variable
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	// Fall back to reading from Pulumi configuration
	cfg := config.New(ctx, "")
	if value := cfg.Get(key); value != "" {
		return value
	}

	return ""
}
