package database

import ()

/*
 * ModelVersion represents an a name and version number for individual model tables.
 * This will allow us to create migrations more effectively.
 */
type ModelVersion struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	TableName string `json:"tablename"`
	Version   int    `json:"version"`
}

/*
 * CurrentVersion takes a model name as a string, and checks the database to see which version of the table
 * is currently in the database. This will allow the migration code to decide what steps need to be taken in order
 * to bring the table to the desired state.
 */
func GetCurrentVersion(model string) int {

	return -1
}
