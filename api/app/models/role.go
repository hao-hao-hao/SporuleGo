package models

import (
	"gopkg.in/mgo.v2/bson"
)

//Role is for permission management
type Role struct {
	ID   bson.ObjectId `bson:"ID"`
	Name string        `bson:"name"`
}
