package umongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func NewDbClient(connectionString string, timeout time.Duration) (*mongo.Client, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		return nil, cancel, err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, cancel, err
	}
	return client, cancel, nil
}
