package app

import (
	"database/sql"
	"time"
	"yahfiilham/go-rest-api/helper"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "yahfiilham@tcp(localhost:3360)/go_restful_api")
	helper.PanicIfError(err)

	db.SetConnMaxIdleTime(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(60 *time.Minute)
	db.SetConnMaxLifetime(60 *time.Minute)

	return db
}