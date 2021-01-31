package main

import (
	"fmt"
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

	err := database.InsertObject("testobjects", myValue)
	if err != nil {
		fmt.Println("Didn't insert test object.")
		panic(err)
	}

	log.Fatal(http.ListenAndServe("localhost:8989", config.Config.Router))
}
