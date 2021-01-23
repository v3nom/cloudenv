package cloudenv

import (
	"context"
	"testing"
)

func TestIsDevContext(t *testing.T) {
	ctx := context.TODO()
	isdevServer := IsDevContext(ctx)
	if isdevServer {
		t.Fatalf("Expected non-dev environment")
	}
}
