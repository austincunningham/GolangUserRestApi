package dao

import (
	"log"

	//"GolangUserRestApi/models"
	mgo "gopkg.in/mgo.v2"
	_ "gopkg.in/mgo.v2/bson"
)

type UsersDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "users"
)

func (m *UsersDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

// func (m *UsersDAO) FindAll() ([]User, error){
// 	var users []User
// 	err := db.C(COLLECTION).Find(bson.M{}).All(&users)
// 	return users, err
// }

// func (m *UsersDAO) Insert(user User) error {
// 	err := db.C(COLLECTION).Insert(&user)
// 	return err
// }