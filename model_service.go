package umongo

import (
	"go.mongodb.org/mongo-driver/mongo"
)

// DataStructure for storing information for a collection and a handle for modifying it
// A convenience function for creating indexes is provided
type ModelService struct {
	Database   string
	Collection string
	Col        *mongo.Collection
}

func NewModelService(client *mongo.Client, database string, collection string) ModelService {
	return ModelService{
		Database:   database,
		Collection: collection,
		Col:        client.Database(database).Collection(collection),
	}
}
