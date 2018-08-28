# Golang REST API

User REST API with crudl operation on mongo db 

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