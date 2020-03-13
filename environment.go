package cloudenv

import (
	"context"
	"log"
	"os"

	cloudtasks "cloud.google.com/go/cloudtasks/apiv2"
	"cloud.google.com/go/compute/metadata"
	"cloud.google.com/go/datastore"
	"cloud.google.com/go/logging"
	"github.com/v3nom/logas"
	"github.com/v3nom/pipes"
)

var isDevAppServer bool
var projectID string
var locationID string

// Clients
var datastoreClient *datastore.Client
var cloudtasksClient *cloudtasks.Client
var loggingClient *logging.Client

// Config config
type Config struct {
	Datastore  bool
	Logging    bool
	CloudTasks bool
}

// Init init
func Init(config Config) map[pipes.ContextKey]interface{} {
	isDevAppServer = os.Getenv("IS_DEVELOPMENT") == "true"
	loadMetadata()

	appContext := map[pipes.ContextKey]interface{}{
		IsDevServer: isDevAppServer,
	}

	// Clients
	ctx := context.Background()
	if config.Datastore {
		initDatastore(ctx)
		appContext[DatastoreClient] = datastoreClient
	}
	if config.Logging && projectID != "" {
		initLogging(ctx)
		appContext[logas.LoggingClient] = loggingClient
	}
	if config.CloudTasks {
		initCloudTasks(ctx)
		appContext[CloudtasksClient] = cloudtasksClient
	}

	return appContext
}

// IsDev is dev
func IsDev() bool {
	return isDevAppServer
}

func initDatastore(ctx context.Context) {
	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("datastore.NewClient. Err: %v", err)
		return
	}
	datastoreClient = client
}

func initLogging(ctx context.Context) {
	client, err := logging.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("logging.NewClient. Err: %v", err)
		return
	}
	loggingClient = client
}

func initCloudTasks(ctx context.Context) {
	client, err := cloudtasks.NewClient(ctx)
	if err != nil {
		log.Fatalf("cloudtasks.NewClient. Err: %v", err)
		return
	}
	cloudtasksClient = client
}

func loadMetadata() {
	if metadata.OnGCE() {
		pid, err := metadata.ProjectID()
		zid, err := metadata.Zone()
		log.Printf("Project: %v, %v", pid, err)
		log.Printf("Zone: %v, %v", zid, err)
		projectID = pid
		locationID = zid
	} else {
		log.Println("Not on GCE")
	}
}

// Dispose disposes resources
func Dispose() {
	if loggingClient != nil {
		loggingClient.Close()
	}
}
