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

func getItems(w http.ResponseWriter, r *http.Request) {

	queryValues := r.URL.Query()
    id_str := queryValues.Get("id")

    // Convert the id parameter to an integer
    id, err := strconv.Atoi(id_str)
    if err != nil && id_str != "" { // does not execute if no id_str has been passed
        http.Error(w, "Invalid id parameter", http.StatusBadRequest)
        return
    }

	var items []Item
	if id == 0 { // means id was not specified
		items = getAllItems()
	} else {
		items = getItemById(id)
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

	if item.ID == 0 || item.Nome == "" || item.Descricao == "" || item.Preco == 0 {
		fmt.Fprintf(w,"One or more invalid parameters")
		return
	}

	if createItem_db(item) {
		fmt.Fprintf(w,"Item successfully created")
	} else {
		fmt.Fprintf(w,"Failed to create item on database")
	}
}

func deleteItem(w http.ResponseWriter, r *http.Request) {

	var item Item
	json.NewDecoder(r.Body).Decode(&item)
			
	if item.ID == 0 {
		fmt.Fprintf(w,"Invalid item id")
		return
	}

	if deleteItem_db(item.ID) {
		w.WriteHeader(http.StatusOK)
		//fmt.Fprintf(w,"Item successfully deleted")
	} else {
		fmt.Fprintf(w,"Failed to delete item from database")
	}
}

func updateItem(w http.ResponseWriter, r *http.Request) {
	var item Item
	json.NewDecoder(r.Body).Decode(&item)

	if item.ID == 0 {
		fmt.Fprintf(w,"Invalid item id")
		return
	}

	oldItems := getItemById(item.ID)
	if oldItems == nil {
        http.Error(w, "No element found to update", http.StatusBadRequest)
        return			
	}

	if updateItem_db(item) {
		fmt.Fprintf(w,"Item successfully updated")
	} else {
		fmt.Fprintf(w,"Failed to update item on database")
	}

}