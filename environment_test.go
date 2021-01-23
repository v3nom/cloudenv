package cloudenv

import (
	"context"
	"testing"
)

func TestDefault(t *testing.T) {
	isGCEOverride = false

	ctx := context.Background()
	err := Init(ctx, Config{
		Datastore:  false,
		CloudTasks: false,
	})

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if !IsDev() {
		t.Fatalf("Expected dev environment")
	}
}

func TestDataStoreClient(t *testing.T) {
	isGCEOverride = false

	ctx := context.Background()

	projectID = "test"
	err := Init(ctx, Config{
		Datastore:  true,
		CloudTasks: false,
	})
	projectID = ""

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	defer func() {
		Dispose()
		if datastoreClient != nil {
			t.Fatalf("Expected client to be closed")
		}
	}()

	requestContext := context.Background()
	requestContext = AddContextValues(requestContext)

	values := ContextValues()
	if values[datastoreClientKey] == nil {
		t.Fatalf("Expected client to be not nil")
	}

	client := Datastore(requestContext)
	if client == nil {
		t.Fatalf("Expected client to be created")
	}
}

func TestError(t *testing.T) {
	isGCEOverride = false

	projectID = ""
	ctx := context.Background()
	err := Init(ctx, Config{
		Datastore:  true,
		CloudTasks: false,
	})
	if err == nil {
		t.Fatalf("Expected an error")
	}
}
