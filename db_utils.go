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

func getItems() []Item {
    
	rows, err := db.Query("SELECT * FROM golang_api")
    if err != nil {
        log.Fatal(err)
        defer rows.Close()
    }

    var items []Item
    for rows.Next() {
        var item Item
        if err := rows.Scan(&item.ID, &item.Nome, &item.Descricao, &item.Preco); err != nil {
            log.Fatal(err)
        }
        items = append(items, item)
    }

    return items
}

func getItemById(id int) []Item {

    rows, err := db.Query("SELECT * FROM golang_api WHERE ID=? LIMIT 1",id)
    if err != nil {
        log.Fatal(err)
        defer rows.Close()
    }

    var items []Item
    for rows.Next() {
        var item Item
        if err := rows.Scan(&item.ID, &item.Nome, &item.Descricao, &item.Preco); err != nil {
            log.Fatal(err)
        }
        items = append(items, item)
    }

    return items
}

func createItem(item Item) bool {

    _, err := db.Exec("INSERT INTO golang_api (id,nome,descricao,preco) VALUES (?,?,?,?)", item.ID, item.Nome, item.Descricao, item.Preco)
    if err != nil {
        log.Fatal(err)
        return false
    }
    return true
}

func deleteItem(id int) bool {

    _, err := db.Exec("DELETE FROM golang_api WHERE id=?",id)
    if err != nil {
        log.Fatal(err)
        return false
    }
    return true
}