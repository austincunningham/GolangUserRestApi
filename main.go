package main

import (
	"fmt"
	"log"
	"net/http"
	"context"
	"encoding/json"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/gorilla/mux"
    "github.com/austincunningham/GolangUserRestApi/models"
)

type DB struct {
	collection *mongo.Collection
}

// find all users
func (db *DB)AllUsers(res http.ResponseWriter, req *http.Request){
	fmt.Println("AllUsers GET")
	// create an array of users
	var results []models.User
	var user models.User
	// set the api header
	res.Header().Set("content-type", "application/json")
	// set the find options, not sure I need this
	findOptions := options.Find()
	// use the find command to get all
	result , err := db.collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		fmt.Println("AllUsers GET failed to query DB", err)
	}
	//go through the result and decode each element at a time
	for result.Next(context.TODO()){
		err := result.Decode(&user)
        if err != nil {
            log.Fatal(err)
		}
		// add to the array
        results = append(results, user)
	}
	//return the array as json
	json.NewEncoder(res).Encode(results)
}

// find a single user
func (db *DB)FindUser(res http.ResponseWriter, req *http.Request){
	fmt.Println("FindUser GET")
	var user models.User
	params := mux.Vars(req)
	objectId, _ := primitive.ObjectIDFromHex(params["id"])
	res.Header().Set("content-type", "application/json")
	filter := bson.M{"_id": objectId}
	err := db.collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil{
		fmt.Println("error",err)
	}
	json.NewEncoder(res).Encode(user)

}

func (db *DB)CreateUser(res http.ResponseWriter, req *http.Request){
	fmt.Println("CreateUser POST")
    var user models.User
	res.Header().Set("content-type", "application/json")
	_ = json.NewDecoder(req.Body).Decode(&user)

	result, err := db.collection.InsertOne(context.TODO(), user)
	if err != nil {
		fmt.Println("CreateUser Error inserting record ", err)
	}
	json.NewEncoder(res).Encode(result)
}

func (db *DB)UpdateUser(res http.ResponseWriter, req *http.Request){
	fmt.Println("UpdateUser PUT")
	var user models.User
	// get the id from the url
	params := mux.Vars(req)
	objectId, _ := primitive.ObjectIDFromHex(params["id"])
	// set the header info
	res.Header().Set("content-type", "application/json")
	//set the filter on the id
	filter := bson.M{"_id": objectId}
	// decode the request body 
	_ = json.NewDecoder(req.Body).Decode(&user)
	update := bson.M{"$set": &user}
	result,err := db.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil{
		fmt.Println("UpdateOne PUT failed to query DB ",err)
	}
	json.NewEncoder(res).Encode(result)
}

func (db *DB)DeleteUser(res http.ResponseWriter, req *http.Request){
	fmt.Println("DeleteUser DELETE")
	params := mux.Vars(req)
	objectId, _ := primitive.ObjectIDFromHex(params["id"])
	res.Header().Set("content-type", "application/json")
	filter := bson.M{"_id": objectId}
	result, err := db.collection.DeleteOne(context.TODO(), filter)
	if err != nil{
		fmt.Println("DeleteUser DELETE failed to query DB",err)
	}
	json.NewEncoder(res).Encode(result)
}


// Define the routes
func main() {
	fmt.Printf("REST API User from golang\n")

	// connect to mongodb
	// Set client otions
    clientOptions := options.Client().ApplyURI("mongodb://admin:admin@mongodb:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())

	// set the collection and database
	collection := client.Database("golang-user").Collection("users")
	// you can now update the global db with collection
	db := &DB{collection: collection }
	

    
	fmt.Println("Connected to MongoDB!")

	//outputs
	fmt.Printf("Server listing on http://mongo:8080/users")
	fmt.Printf("\nCTRL C to exit\n")

	// Controller for endpoints
	r := mux.NewRouter()
	r.HandleFunc("/users", db.AllUsers).Methods("GET")
	r.HandleFunc("/users/{id}", db.FindUser).Methods("GET")
	r.HandleFunc("/users", db.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", db.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", db.DeleteUser).Methods("DELETE")

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
