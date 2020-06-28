package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// the properties in mongodb document
type User struct {
	Id primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string `bson:"name" json:"name"`
	Email string `bson:"email" json:"email"`
	Password string `bson:"password" json:"password"`
}