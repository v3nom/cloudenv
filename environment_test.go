package cloudenv

import "testing"

func TestDefault(t *testing.T) {
	Init(Config{
		Datastore:  false,
		CloudTasks: false,
		Logging:    false,
	})

	if IsDev() {
		t.Fatalf("Expected non-dev environment")
	}
}
