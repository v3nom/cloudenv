package cloudenv

import (
	"context"
)

// Config config
type Config struct {
	Datastore  bool
	CloudTasks bool
}

type contextKey string

var appContext = map[contextKey]interface{}{}

// Init initializes cloud env
func Init(ctx context.Context, config Config) error {
	err := loadMetadata()
	if err != nil {
		return err
	}

	appContext[isDevServerKey] = IsDev()

	// Clients
	if config.Datastore {
		err := initDatastore(ctx)
		if err != nil {
			return err
		}
		appContext[datastoreClientKey] = datastoreClient
	}

	if config.CloudTasks {
		err := initCloudTasks(ctx)
		if err != nil {
			return err
		}
		appContext[cloudtasksClientKey] = cloudtasksClient
	}

	return nil
}

// Dispose disposes resources
func Dispose() {
	if datastoreClient != nil {
		datastoreClient.Close()
		datastoreClient = nil
	}
	if cloudtasksClient != nil {
		cloudtasksClient.Close()
		cloudtasksClient = nil
	}
}

// AddContextValues adds context values
func AddContextValues(ctx context.Context) context.Context {
	for k, v := range appContext {
		ctx = context.WithValue(ctx, k, v)
	}
	return ctx
}

// ContextValues context values
func ContextValues() map[interface{}]interface{} {
	values := map[interface{}]interface{}{}
	for k, v := range appContext {
		values[k] = v
	}
	return values
}
