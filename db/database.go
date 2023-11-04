// Package db provides access to the Datastore API.
package db

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/datastore"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
)

type DatabaseThingy struct {
	client *datastore.Client
	ctx    context.Context
}

func NewDatabaseThingy() (*DatabaseThingy, error) {
	ctx := context.Background()
	projectID := getEnvVar("DATASTOREDB_PROJECT")
	devMode := getEnvVar("DEVMODE") == "true"

	var client *datastore.Client
	var err error

	if devMode {
		host := getEnvVar("DATASTOREDB_HOST")
		port := getEnvVar("DATASTOREDB_PORT")
		endpoint := fmt.Sprintf("%s:%s", host, port)
		client, err = datastore.NewClient(ctx, projectID, option.WithEndpoint(endpoint), option.WithGRPCDialOption(grpc.WithInsecure()))
	} else {
		client, err = datastore.NewClient(ctx, projectID)
	}

	if err != nil {
		return nil, err
	}

	return &DatabaseThingy{
		client: client,
		ctx:    ctx,
	}, nil
}

func getEnvVar(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return ""
}

// Select retrieves entities from Datastore. The kind must match the kind of entities stored in Datastore.
func (d *DatabaseThingy) Select(kind string, dst interface{}) error {
	query := datastore.NewQuery(kind)
	_, err := d.client.GetAll(d.ctx, query, dst)
	return err
}

// Insert adds a new entity to Datastore.
func (d *DatabaseThingy) Insert(kind string, src interface{}) (*datastore.Key, error) {
	key := datastore.IncompleteKey(kind, nil)
	return d.client.Put(d.ctx, key, src)
}
