package database

import (
	"database/sql"
)

var Tables *sql.DB

// For open connect in db postgreSQL
const (
	user    string = "postgres"
	pass    string = "537537"     //aezakmi202
	dbname  string = "main_inote" // main_inote
	sslmode string = "disable"
)

// Formation full string for authorization
const ConnectString string = "user=" + user + " password=" + pass + " dbname=" + dbname + " sslmode=" + sslmode

// DATABASE QUERIES
// Requests internal/controllers
const SelectUserData string = `SELECT ud.id, ud.user_custom_id FROM users_data ud WHERE ud.token=$1`

// Requests api/autorization/signin
const SelectUserDataSignIn string = `SELECT ud.id, u.network_status, ud.token FROM users u, users_data ud WHERE ud.login=$1 AND ud.password=$2 AND ud.id=u.id`

// Requests api/autorization/signout
const SelectUserIdForToken string = `SELECT ud.id FROM users_data ud WHERE token=$1`

// Requests api/autorization/signup
const (
	SelectLogin       string = `SELECT login FORM users_data WHERE login=$1`
	InsertNewUserData string = `INSERT INTO users_data (login, password, email, token) VALUES ($1, $2, $3, $4)`
	InsertNewUser     string = `INSERT INTO users (name) VALUES ($1)`
)

// Requests api/checkOnline
const (
	SelectLoginForIdToken string = `SELECT login FROM users_data WHERE id=$1 AND token=$2`
)

// Requests api/userData
const (
	SelectIdForIdOrCustomId string = `SELECT id FROM users_data WHERE id=$1 OR user_custom_id=$2`
	SelectProfileData       string = `SELECT logo, banner, name, position, audience, verification, network_status FROM users WHERE id=$1`
)

// Update user Token
const (
	UpdateTokenNotNull string = `UPDATE users_data SET token=$1 WHERE users_data.id=$2`
	UpdateTokenNull    string = `UPDATE users_data SET token=null WHERE id=$1`
)

// Update user Network Status
const (
	UpdateNetworkStatusOnline  string = `UPDATE users SET network_status='online' WHERE id=$1 AND network_status<>'online'`
	UpdateNetworkStatusOffline string = `UPDATE users SET network_status='offline' WHERE id=$1 AND network_status<>'offline'`
)
