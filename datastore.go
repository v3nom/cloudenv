package cloudenv

import (
	"context"

	"cloud.google.com/go/datastore"
)

var datastoreClientKey = contextKey("CloudEnv_datastoreClient")
var datastoreClient *datastore.Client

func initDatastore(ctx context.Context) error {
	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		return err
	}
	datastoreClient = client

	return nil
}

// Datastore gets datastore client from context
func Datastore(ctx context.Context) *datastore.Client {
	client := ctx.Value(datastoreClientKey)
	if client == nil {
		return nil
	}
	return client.(*datastore.Client)
}
