package config

import (
	"context"
	CagoxConfig "github.com/cagox/config"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"os"
)

//Config is the configuration variable for the application
var Config *OCFConfigStruct

type OCFConfigStruct struct {
	CagoxConfig.ConfigurationStruct

	SiteName string
	LogPath  string

	//The items below are not in the JSON file.
	Router       *mux.Router
	Logger       *log.Logger
	MongoClient  *mongo.Client
	MongoContext context.Context
	LogFile      *os.File
}

func init() {
	Config = &OCFConfigStruct{}
	loadConfigs()
	Config.Router = mux.NewRouter()
}

func loadConfigs() {
	CagoxConfig.LoadConfigs(Config, "OCFCONF")
}
