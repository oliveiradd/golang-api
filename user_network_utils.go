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

func getUser(w http.ResponseWriter, r *http.Request) bool {

	var loginCredentials LoginCredentials
    json.NewDecoder(r.Body).Decode(&loginCredentials)

    if len(loginCredentials.Email) > 255 {
		http.Error(w, "Invalid e-mail address", http.StatusBadRequest)
		return
	}

	if len(loginCredentials.PWHash) > 255 {
		http.Error(w,"Invalid password hash", http.StatusBadRequest)
		return
	}
	
	user := getUser()

	if user == nil {
		http.Error(w,"User does not exist", http.StatusNotFound)
	}

	if user.PWHash != loginCredentials.PWHash {
		http.Error(w,"Invalid password",http.StatusUnauthorized)
	}

	jsonData, err := json.Marshal(user)
    w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(StatusOK)
	w.Write(jsonData)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
    json.NewDecoder(r.Body).Decode(&user)
	// continuar daqui 
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