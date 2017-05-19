package model

import (
    "gopkg.in/mgo.v2/bson"
)

type City struct {
    ID      bson.ObjectId   `json:"id" bson:"_id,omitempty"`
    Name    string          `json:"name" bson:"name"`
    Country string          `json:"country" bson:"country"`
}