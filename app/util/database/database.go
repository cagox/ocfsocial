package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"

	"github.com/cagox/ocfsocial/app/util/config"
)

//DialMongoSession starts the main mongo session.
func DialMongoSession() {
	credential := options.Credential{
		AuthMechanism: "SCRAM-SHA-1",
		Username:      config.Config.DatabaseUserName,
		Password:      config.Config.DatabasePassword,
		AuthSource:    config.Config.DatabaseName,
	}

	clientOpts := options.Client().ApplyURI(config.Config.DatabaseServerURL).SetAuth(credential)

	client, err := mongo.Connect(config.Config.MongoContext, clientOpts)
	if err != nil {
		panic(err)
	}

	config.Config.MongoClient = client

	err = client.Ping(config.Config.MongoContext, nil)
	if err != nil {
		fmt.Println("Database didn't connect at all.")
		log.Fatal(err)
	}

}

func buildContext() context.Context {
	return context.TODO()
}

func init() {
	config.Config.MongoContext = buildContext()

}
