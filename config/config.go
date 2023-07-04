package config

import (
	"database/sql"
)

func GetMySQLDB() (db *sql.DB, err error) {

	dbDriver := "mysql"
	dbUser := "daniel"
	dbPass := "oculos18"
	dbName := "younit"
	// db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName+"?parseTime=true")
	// The parseTime=true is to enable TIMESTAMP into Time.Time Golang

	return
}
