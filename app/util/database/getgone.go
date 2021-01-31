package database

import (
	"github.com/cagox/ocfsocial/app/util/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

//Get's one object from the database.
func GetOne(collectionName string, searchQuery bson.D, recipient interface{}) {
	collection := config.Config.MongoClient.Database(config.Config.DatabaseName).Collection(collectionName)

	err := collection.FindOne(config.Config.MongoContext, searchQuery).Decode(&recipient)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			recipient = nil
		} else {
			log.Fatal(err)
		}
	}

}
