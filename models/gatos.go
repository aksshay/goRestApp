package models

import (
  "gopkg.in/mgo.v2/bson"
)

type Gato struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Name        string        `bson:"name" json:"name"`
	Image       string        `bson:"image" json:"image"`
	Description string        `bson:"description" json:"description"`
}
