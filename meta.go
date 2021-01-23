package cloudenv

import (
	"context"

	"cloud.google.com/go/compute/metadata"
)

var isDevServerKey = contextKey("CloudEnv_isdev")

var isDevAppServer bool
var projectID string
var zoneID string
var defaultServiceAccountEmail string

// IsDev is dev server
func IsDev() bool {
	return isDevAppServer
}

// IsDevContext is dev server
func IsDevContext(ctx context.Context) bool {
	isDev := ctx.Value(isDevServerKey)
	if isDev == nil {
		return false
	}
	return ctx.Value(isDevServerKey).(bool)
}

// ProjectID project ID
func ProjectID() string {
	return projectID
}

// ZoneID zone ID
func ZoneID() string {
	return zoneID
}

// DefaultEmail default service account email
func DefaultEmail() string {
	return defaultServiceAccountEmail
}

func loadMetadata() error {
	if !metadata.OnGCE() {
		isDevAppServer = true
		return nil
	}

	isDevAppServer = false

	// Project
	pid, err := metadata.ProjectID()
	if err != nil {
		return err
	}
	projectID = pid

	// Zone
	zid, err := metadata.Zone()
	if err != nil {
		return err
	}
	zoneID = zid

	// Service account email
	email, err := metadata.Email("default")
	if err != nil {
		return err
	}
	defaultServiceAccountEmail = email

	return nil
}
