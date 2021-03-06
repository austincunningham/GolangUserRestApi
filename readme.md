# Golang REST API

User REST API with crudl operation on mongo db. Following endpoints will be available 

    /users GET         (get all users)
    /users/{id}        (find a single user by id)
    /users POST        (create a user)
    /users/{id} PUT    (update a user)
    /users/{id} DELETE (delete a user) 

## Prerequisites
- Mongo db
- Golang
Install the following
```bash
go get go.mongodb.org/mongo-driver
go get github.com/gorilla/mux  
```

## Usage
Use the following command

```bash
go run main.go
```