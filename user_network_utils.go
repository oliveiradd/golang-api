package main

import (
//    "database/sql"
    "encoding/json"
    "fmt"
//    "log"
    "net/http"
//    "strconv"


    _ "github.com/go-sql-driver/mysql"
//    "github.com/gorilla/mux"
)

func getUser(w http.ResponseWriter, r *http.Request) {

	var loginCredentials LoginCredentials
    json.NewDecoder(r.Body).Decode(&loginCredentials)

    if len(loginCredentials.Email) > 255 {
		http.Error(w, "Invalid e-mail address", http.StatusBadRequest)
		return
	}

	if len(loginCredentials.PWHash) > 255 || loginCredentials.PWHash == "" {
		http.Error(w,"Invalid password hash", http.StatusBadRequest)
		return
	}
	
	user := getUserByEmail(loginCredentials.Email)

	if user == nil {
		http.Error(w,"User does not exist", http.StatusNotFound)
		return
	}

	if user[0].PWHash != loginCredentials.PWHash {
		http.Error(w,"Invalid password",http.StatusUnauthorized)
		return
	}

	jsonData, err := json.Marshal(user)
    if err != nil {
		http.Error(w,"Failed to get json from user",http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
	return
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
    json.NewDecoder(r.Body).Decode(&user)
	// continuar daqui 
	if user.Email == "" || user.PWHash == "" || user.Name == "" || user.CreatedAt == "" {
		fmt.Fprintf(w,"One or more invalid parameters")
		return
	}
	
	if createUser_db(user) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w,"User successfully created")
	} else {
		http.Error(w,"Failed to create user on database",http.StatusInternalServerError)
	}
}

func deleteUser(w http.ResponseWriter, r *http.Request) {

	var loginCredentials LoginCredentials
    json.NewDecoder(r.Body).Decode(&loginCredentials)
	
	if getUserByEmail(loginCredentials.Email) != nil {
		return
	}

	if deleteUser_db(loginCredentials.Email) {
		w.WriteHeader(http.StatusOK)
	} else {

		fmt.Fprintf(w,"Failed to delete user from database")
	}

}

func updateUser(w http.ResponseWriter, r *http.Request) {
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