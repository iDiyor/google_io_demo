package model

import (
    "gopkg.in/mgo.v2/bson"
)

type Venue struct {
	ID          bson.ObjectId `json:"id" bson:"_id,omitempty"`
	CityID      string 		  `json:"city_id" bson:"city_id"`
	Name        string 		  `json:"name" bson:"name"`
	Description string 		  `json:"description" bson:"description"`
    Address     string 		  `json:"address" bson:"address"`
    ImageURL    string 		  `json:"image_url" bson:"image_url"`
}
