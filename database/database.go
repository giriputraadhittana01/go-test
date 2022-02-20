package database

import "fmt"

var connection string

func init() {
	connection = "MSSQL"
}

func GetConnection() string {
	return connection
}
