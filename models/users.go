package models

import "gopkg.in/mgo.v2/bson"

// Represents a user, the bson keyword tells the mgo driver how to name
// the properties in mongodb document
type User struct {
	ID bson.ObjectId `bson:"_id" json:"id"`
	Name string `bson:"name" json:"name"`
	Email string `bson:"email" json:"email"`
	Password string `bson:"password" json:"password"`
}