package umongo

import (
	"context"
	"fmt"

	"github.com/dunv/ulog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// DataStructure for storing information for a collection and a handle for modifying it
// A convenience function for creating indexes is provided
type ModelService struct {
	Database         string
	Collection       string
	CollectionHandle *mongo.Collection
	Indexes          []mongo.IndexModel
}

func NewModelService(client *mongo.Client, database string, collection string, indexes []mongo.IndexModel) ModelService {
	return ModelService{
		Database:         database,
		Collection:       collection,
		CollectionHandle: client.Database(database).Collection(collection),
		Indexes:          indexes,
	}
}

// Identifies existing indexes by name (and only by name!) and adds desired indexes if
// their name does not exist in the existing indexes
// Options, fields, etc. are not regarded
// It seems the mongo-driver already handles duplicate indexes, it returns OK and does not create new ones
func (s ModelService) EnsureIndexes() error {
	if s.Indexes != nil {
		// Load existing indexes
		ctx := context.Background()
		indexView := s.CollectionHandle.Indexes()
		cur, err := indexView.List(ctx)
		if err != nil {
			return err
		}
		existingIndexes := []bson.D{}
		for cur.Next(ctx) {
			index := bson.D{}
			if err := cur.Decode(&index); err != nil {
				return fmt.Errorf("unable to decode bson index document (%v)", err)
			}
			existingIndexes = append(existingIndexes, index)
		}

		for _, existingIndex := range existingIndexes {
			if existingIndexName, ok := existingIndex.Map()["name"]; ok {
				ulog.Infof("Found index on db: %s, collection: %s, name: %s", s.Database, s.Collection, existingIndexName)
			}
		}

		// Check if desiredIndex is already there
		nonExistingIndexes := []mongo.IndexModel{}
		for _, desiredIndex := range s.Indexes {
			found := false
			for _, existingIndex := range existingIndexes {
				if existingIndexName, ok := existingIndex.Map()["name"]; ok {
					if desiredIndex.Options.Name != nil && *desiredIndex.Options.Name == existingIndexName {
						found = true
					}
				}
			}

			if !found {
				nonExistingIndexes = append(nonExistingIndexes, desiredIndex)
			}
		}

		errors := []error{}
		for _, nonExistingIndex := range nonExistingIndexes {
			res, err := indexView.CreateOne(ctx, nonExistingIndex)
			if err != nil {
				errors = append(errors, err)
			} else {
				ulog.Infof("Created missing index on db: %s, collection: %s, name: %s, res: %s", s.Database, s.Collection, *nonExistingIndex.Options.Name, res)
			}

		}
		if len(errors) != 0 {
			return fmt.Errorf("err creating indexes (%v)", errors)
		}
	}

	return nil
}
