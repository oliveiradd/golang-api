package main

import (
    "database/sql"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "strconv"


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

func getItems(w http.ResponseWriter, r *http.Request) {
    
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

    jsonData, err := json.Marshal(items)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(jsonData)
}

func getItemById(w http.ResponseWriter, r *http.Request) {

	queryValues := r.URL.Query()
    id_str := queryValues.Get("id")

    // Convert the id parameter to an integer
    id, err := strconv.Atoi(id_str)
    if err != nil {
        http.Error(w, "Invalid id parameter", http.StatusBadRequest)
        return
    }

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

    jsonData, err := json.Marshal(items)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(jsonData)
}

func createItem(w http.ResponseWriter, r *http.Request) {

    var item Item
    json.NewDecoder(r.Body).Decode(&item)

    _, err := db.Exec("INSERT INTO golang_api (id,nome,descricao,preco) VALUES (?,?,?,?)", item.ID, item.Nome, item.Descricao, item.Preco)
    if err != nil {
        log.Fatal(err)
    }
    w.WriteHeader(http.StatusCreated)
    fmt.Fprintf(w, "Item successfully created")
}