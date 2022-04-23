package database

import (
	"database/sql"
	"fmt"
)

var Tables *sql.DB

// Constants for authorization in db postgreSQL
const (
	user    string = "postgres"
	pass    string = "537537"     //aezakmi202
	dbname  string = "main_inote" // main_inote
	sslmode string = "disable"
)

// Formation full string for authorization
var ConnectString string = fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", user, pass, dbname, sslmode)
