# Golang REST API

User REST API with crudl operation on mongo db. Following endpoints will be available 

    /users GET    (get all users)
    /users/{id}   (find a single user by id)
    /users POST   (create a user)
    /users PUT    (update a user)
    /users DELETE (delete a user) 

## Prerequisites

- Golang
Install the following
- toml : Parses the configuration file for MongoDB server & credentials
- mux : Request router 
- mgo : MongoDB driver
```bash
go get github.com/BurntSushi/toml
go get gopkg.in/mgo.v2
go get github.com/gorilla/mux  
```

## Usage
Use the following command

```bash
go run app.go
```