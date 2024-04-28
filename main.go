package main

import (
//    "database/sql"
//    "encoding/json"
    "fmt"
    "log"
    "net/http"
//    "strconv"


    _ "github.com/go-sql-driver/mysql"
    "github.com/gorilla/mux"
)


type Item struct {
    ID        int     `json:"id"`
    Nome      string  `json:"nome"`
    Descricao string  `json:"descricao"`
    Preco     float64 `json:"preco"`
}

func main() {
	initDB()
	defer db.Close()

	// initialize router
	router := mux.NewRouter()
    router.HandleFunc("/items", retrieveItems).Methods("GET")
    router.HandleFunc("/items", receiveItem).Methods("POST")
    router.HandleFunc("/item", retrieveItemById).Methods("GET")
    router.HandleFunc("/item", updateItem).Methods("POST")
    
	// start http server
	fmt.Println("Server listening on port 8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}