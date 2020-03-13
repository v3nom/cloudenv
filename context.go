package cloudenv

import "github.com/v3nom/pipes"

// DatastoreClient datastore client
var DatastoreClient pipes.ContextKey = "datastore"

// CloudtasksClient cloudtasks client
var CloudtasksClient pipes.ContextKey = "cloudtasks"

// IsDevServer true if development server
var IsDevServer pipes.ContextKey = "isdev"
