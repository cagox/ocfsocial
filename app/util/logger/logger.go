package logger

import (
	"github.com/cagox/ocfsocial/app/util/config"
	"log"
	"os"
)

//StartLogging() starts the logger.
func StartLogging() {
	//Set up logging
	var err error
	config.Config.LogFile, err = os.OpenFile(config.Config.LogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
		panic("Couldn't Open Log File")
	}
	config.Config.Logger = log.New(config.Config.LogFile, "ocfsocial:", log.LstdFlags)
}

//StopLogging() stops the logger.
func StopLogging() {
	_ = config.Config.LogFile.Close()
	//TODO: Add error handling.
}
