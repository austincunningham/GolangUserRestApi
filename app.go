package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"GolangUserRestApi/config"
	"GolangUserRestApi/dao"
	"GolangUserRestApi/models"
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


//error
// func respondWithError(w http.ResponseWriter, code int, msg string) {
// 	respondWithJson(w, code, map[string]string{"error": msg})
// }
// normal response
// func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
// 	response, _ := json.Marshal(payload)
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(code)
// 	w.Write(response)
// }



// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}


// Define the routes
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
