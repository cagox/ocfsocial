package user

import (
	"github.com/cagox/ocfsocial/app/util/database"
)

func MigrateUp() {
	MigrationsUp := []func(){
		initialMigration,
	}

	dbVersion := database.GetCurrentVersion("user")

	if dbVersion == Version {
		return
	}

	if dbVersion <= 0 {
		for _, funcP := range MigrationsUp {
			funcP()
		}
	}

}

func initialMigration() {

}
