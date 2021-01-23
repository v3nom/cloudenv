package cloudenv

import (
	"context"

	cloudtasks "cloud.google.com/go/cloudtasks/apiv2"
)

var cloudtasksClientKey = contextKey("CloudEnv_cloudtasksclient")
var cloudtasksClient *cloudtasks.Client

func initCloudTasks(ctx context.Context) error {
	client, err := cloudtasks.NewClient(ctx)
	if err != nil {
		return err
	}
	cloudtasksClient = client
	return nil
}

// CloudTasks gets cloudtasks client from context
func CloudTasks(ctx context.Context) *cloudtasks.Client {
	client := ctx.Value(datastoreClientKey)
	if client == nil {
		return nil
	}
	return client.(*cloudtasks.Client)
}
