package main

import (
	"fmt"
	"github.com/cagox/ocfsocial/app/routes"
	"github.com/cagox/ocfsocial/app/util/config"
	"github.com/cagox/ocfsocial/app/util/database"
	"github.com/cagox/ocfsocial/app/util/logger"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
)

type TestStruct struct {
	Name      string
	TestValue string
}

func main() {
	logger.StartLogging()
	defer logger.StopLogging()

	database.DialMongoSession()

	routes.Routes()

	myValue := TestStruct{}
	database.GetOne("testobjects", bson.D{{"name", "Jim"}}, myValue)
	fmt.Printf("%+v", myValue)

	log.Fatal(http.ListenAndServe("localhost:8989", config.Config.Router))
}
