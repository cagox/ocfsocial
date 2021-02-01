package database

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/cagox/ocfsocial/app/util/config"
)

func OpenDatabase() {
	db, err := sql.Open("mysql", databaseURI())
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	config.Config.Database = db
}

func CloseDatabase() {
	config.Config.Database.Close()
}

//Here we just build the URI and return it. By using a function, we can change it in one place later.
func databaseURI() string {
	uriString := config.Config.DatabaseUserName + ":" + config.Config.DatabasePassword + "@/" + config.Config.DatabaseName + "?" + config.Config.DatabaseOptions
	return uriString
}
