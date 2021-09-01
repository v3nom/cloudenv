package cloudenv

import (
	"context"

	"cloud.google.com/go/logging"
)

var loggingClientKey = contextKey("CloudEnv_loggingClient")
var loggingClient *logging.Client

func initLogging(ctx context.Context) error {
	client, err := logging.NewClient(ctx, projectID)
	if err != nil {
		return err
	}

	loggingClient = client
	return nil
}

// Logging gets logging client from context
func Logging(ctx context.Context) *logging.Client {
	client := ctx.Value(loggingClientKey)
	if client == nil {
		return nil
	}
	return client.(*logging.Client)
}
