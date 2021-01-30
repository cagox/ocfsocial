package main

import (
	"github.com/cagox/ocfsocial/app/config"
	"github.com/cagox/ocfsocial/app/database"
	"github.com/cagox/ocfsocial/app/logger"
	"github.com/cagox/ocfsocial/app/routes"
	"log"
	"net/http"
)

func main() {
	logger.StartLogging()
	defer logger.StopLogging()

	database.DialMongoSession()

	routes.Routes()

	log.Fatal(http.ListenAndServe("localhost:8080", config.Config.Router))
}
