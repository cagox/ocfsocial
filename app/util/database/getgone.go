package database

import (
	"fmt"
	"github.com/cagox/ocfsocial/app/util/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"reflect"
)

//Get's one object from the database.
func GetOne(collectionName string, searchQuery bson.D, recipient interface{}) {
	collection := config.Config.MongoClient.Database(config.Config.DatabaseName).Collection(collectionName)

	fmt.Println("Recipient is type: ", reflect.TypeOf(recipient))

	err := collection.FindOne(config.Config.MongoContext, searchQuery).Decode(recipient)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("Couldn't find it.")
			recipient = nil
		} else {
			log.Fatal(err)
		}
	}

}

//TODO: For now, this method is just cruft. It doesn't work.
