package database

/*
 * This code assumes MongoDB.
 * The interface should be able to apply to most databases however.
 */

import (
	"github.com/cagox/ocfsocial/app/util/config"
)

//InsertObject is used to put things in the database.
func InsertObject(collectionName string, object interface{}) error {
	//TODO: Once this is working, add some error handling.
	collection := config.Config.MongoClient.Database(config.Config.DatabaseName).Collection(collectionName)

	_, err := collection.InsertOne(config.Config.MongoContext, object)
	if err != nil {
		return err
	}
	return nil
}
