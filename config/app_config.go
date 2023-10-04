// config/app_config.go
// This file handles the retrieval of application-specific configuration settings.

package config

import (
	pcfg "github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
	"github.com/usrbinkat/pulumi-kind-go/cert-manager/helper"
	"github.com/usrbinkat/pulumi-kind-go/types"
)

const (
	DefaultAppName    = "zot"
	DefaultAppImage   = "docker.io/library/nginx"
	DefaultAppVersion = "latest"
)

// GetAppConfig fetches application configuration settings.
func GetAppConfig(config *pcfg.Config) types.App {
	return types.App{
		Name:    helper.IfEmpty(config.Get("appName"), DefaultAppName),
		Image:   helper.IfEmpty(config.Get("appImage"), DefaultAppImage),
		Version: helper.IfEmpty(config.Get("appVersion"), DefaultAppVersion),
	}
}
