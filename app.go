package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func AllUsers(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "ToDo")
}

func FindUser(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "ToDo")
}

func CreateUser(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "ToDo")
}

func UpdateUsers(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "ToDo")
}

func DeleteUser(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "ToDo")
}

func main() {
	fmt.Printf("REST API User from golang\n")

	//outputs
	fmt.Printf("Server listing on http://localhost:8080/users")
	fmt.Printf("\nCTRL C to exit\n")

	// Controller for endpoints
	r := mux.NewRouter()
	r.HandleFunc("/users", AllUsers).Methods("GET")
	r.HandleFunc("/users/{id}", FindUser).Methods("GET")
	r.HandleFunc("/users", CreateUser).Methods("POST")
	r.HandleFunc("/users", UpdateUsers).Methods("PUT")
	r.HandleFunc("/users", DeleteUser).Methods("DELETE")

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
