package main

import (
    "database/sql"
//    "encoding/json"
//    "fmt"
    "log"
//    "net/http"
//    "strconv"


    _ "github.com/go-sql-driver/mysql"
//    "github.com/gorilla/mux"
)

var db *sql.DB

func initDB() {
	var err error
	db, err = sql.Open("mysql","test_user:test_password@tcp(10.8.0.1:3306)/test_db")
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxOpenConns(10)
}