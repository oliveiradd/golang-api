package main

import (
//    "database/sql"
    "encoding/json"
    "fmt"
//    "log"
    "net/http"
    "strconv"


    _ "github.com/go-sql-driver/mysql"
//    "github.com/gorilla/mux"
)

func retrieveItems(w http.ResponseWriter, r *http.Request) {

	items := getItems()
	jsonData, err := json.Marshal(items)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(jsonData)
}

func retrieveItemById(w http.ResponseWriter, r *http.Request) {

	queryValues := r.URL.Query()
    id_str := queryValues.Get("id")

    // Convert the id parameter to an integer
    id, err := strconv.Atoi(id_str)
    if err != nil {
        http.Error(w, "Invalid id parameter", http.StatusBadRequest)
        return
    }
	items := getItemById(id)

	jsonData, err := json.Marshal(items)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(jsonData)
}

func receiveItem(w http.ResponseWriter, r *http.Request) {
	var item Item
    json.NewDecoder(r.Body).Decode(&item)

	if item.ID == 0 || item.Nome == "" || item.Descricao == "" || item.Preco == 0 {
		fmt.Fprintf(w,"One or more invalid parameters")
		return
	}

	if createItem(item) {
		fmt.Fprintf(w,"Item successfully created")
	} else {
		fmt.Fprintf(w,"Failed to create item on database")
	}
}