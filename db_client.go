package umongo

import (
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

func NewClient(
	connectionString string,
	appName string,
	connectTimeout time.Duration,
	opts ...applyOption,
) (*mongo.Client, error) {
	usedOptions := &clientOpts{
		requireSuccessfulPing:    nil,
		additionalClientOpts:     []*options.ClientOptions{},
		registerEmbeddedDocument: true,
	}

	for _, opt := range opts {
		opt(usedOptions)
	}

	mongoOpts := options.Client().
		ApplyURI(connectionString).
		SetConnectTimeout(connectTimeout).
		SetAppName(appName)

	if usedOptions.registerEmbeddedDocument {
		// Register mapType as default EmbeddedDocument "marshalInto"-type
		// That way if we specify an interface to be rendered into we do not end up with
		// { "Key": "xxx", "Value": "yyy"} but with { "xxx": "yyy" }
		// https://jira.mongodb.org/browse/GODRIVER-988
		tM := reflect.TypeOf(bson.M{})
		reg := bson.NewRegistry()
		reg.RegisterTypeMapEntry(bson.TypeEmbeddedDocument, tM)
		mongoOpts.SetRegistry(reg)
	}

	c, err := mongo.Connect(append([]*options.ClientOptions{mongoOpts}, usedOptions.additionalClientOpts...)...)
	if err != nil {
		return nil, err
	}

	if usedOptions.requireSuccessfulPing != nil {
		err = c.Ping(*usedOptions.requireSuccessfulPing, readpref.Primary())
		if err != nil {
			return nil, err
		}
	}

	return c, nil
}
