package main

import (
	"github.com/cagox/ocfsocial/app/routes"
	"github.com/cagox/ocfsocial/app/util/config"
	"github.com/cagox/ocfsocial/app/util/database"
	"github.com/cagox/ocfsocial/app/util/logger"
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

	myValue := TestStruct{"Jim", "I like this."}

	database.InsertObject("testobjects", myValue)

	log.Fatal(http.ListenAndServe("localhost:8989", config.Config.Router))
}
