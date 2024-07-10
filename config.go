package main

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

func initDB() {
	var err error
	connStr := "user=postgres password=NjjECNPSj?0jPQWRDLvi dbname=logistics_tutti host=5.78.72.251 sslmode=disable"
	db, err = sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatalln(err)
	}
	db.SetMaxOpenConns(99)
	db.SetMaxIdleConns(99)
}
