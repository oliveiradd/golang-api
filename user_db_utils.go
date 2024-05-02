package main

import (
//    "database/sql"
//    "encoding/json"
//    "fmt"
    "log"
//    "net/http"
//    "strconv"


    _ "github.com/go-sql-driver/mysql"
//    "github.com/gorilla/mux"
)

func getAllUsers() []User {
    
	rows, err := db.Query("SELECT * FROM user_golang_api")
    if err != nil {
        log.Fatal(err)
        defer rows.Close()
    }

    var users []User
    for rows.Next() {
        var user User
        if err := rows.Scan(&user.Email, &user.PWHash, &user.Name, &user.CreatedAt, &user.LastLogin); err != nil {
            log.Fatal(err)
        }
        users = append(users, user)
    }

    return users
}

func getUserByEmail(email string) []User {

    rows, err := db.Query("SELECT * FROM user_golang_api WHERE email=? LIMIT 1",email)
    if err != nil {
        log.Fatal(err)
        defer rows.Close()
    }

    var users []User
    for rows.Next() {
        var user User
        if err := rows.Scan(&user.Email, &user.PWHash, &user.Name, &user.CreatedAt, &user.LastLogin); err != nil {
            log.Fatal(err)
        }
        users = append(users, user)
    }

    return users
}

func createUser_db(user User) bool {

    _, err := db.Exec("INSERT INTO user_golang_api (email,pwhash,name,created_at,last_login) VALUES (?,?,?,?,?)", user.Email, user.PWHash, user.Name, user.CreatedAt, user.LastLogin)
    if err != nil {
        log.Fatal(err)
        return false
    }
    return true
}

func deleteUser_db(email string) bool {

    _, err := db.Exec("DELETE FROM user_golang_api WHERE email=?",email)
    if err != nil {
        log.Fatal(err)
        return false
    }
    return true
}

func updateUser_db(user User) bool {
      
    _, err :=db.Exec("UPDATE user_golang_api SET nome=?,descricao=?,preco=? WHERE id=?", user.Email, user.PWHash, user.Name, user.CreatedAt,user.LastLogin)
    if err != nil {
        log.Fatal(err)
        return false
    }
    return true
}