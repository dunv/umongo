package umongo

import (
	"context"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func NewDbClient(connectionString string, timeout time.Duration) (*mongo.Client, context.CancelFunc, error) {
	// Register mapType as default EmbeddedDocument "marshalInto"-type
	// That way if we specify an interface to be rendered into we do not end up with
	// { "Key": "xxx", "Value": "yyy"} but with { "xxx": "yyy" }
	// https://jira.mongodb.org/browse/GODRIVER-988
	tM := reflect.TypeOf(bson.M{})
	reg := bson.NewRegistryBuilder().RegisterTypeMapEntry(bsontype.EmbeddedDocument, tM).Build()

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	client, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI(connectionString),
		options.Client().SetRegistry(reg),
	)
	if err != nil {
		return nil, cancel, err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, cancel, err
	}
	return client, cancel, nil
}
