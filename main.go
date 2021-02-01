package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/cagox/ocfsocial/app/routes"
	"github.com/cagox/ocfsocial/app/util/config"
	"github.com/cagox/ocfsocial/app/util/database"
	"github.com/cagox/ocfsocial/app/util/logger"
)

func main() {
	logger.StartLogging()
	defer logger.StopLogging()

	database.OpenDatabase()
	defer database.CloseDatabase()

	routes.Routes()

	log.Fatal(http.ListenAndServe("localhost:8989", config.Config.Router))
}

func printOne() {
	fmt.Println("This is the first line.")
}

func printTwo() {
	fmt.Println("This is the second line.")
}

func printThree() {
	fmt.Println("This is the third line.")
}
