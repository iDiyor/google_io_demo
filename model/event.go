package model

import (
    "time"
    "gopkg.in/mgo.v2/bson"
)
    
type Event struct {
    ID          bson.ObjectId `json:"id" bson:"_id,omitempty"`
    VenueID     string        `json:"venue_id" bson:"venue_id"`
    Name        string        `json:"name" bson:"name"`
    Description string        `json:"description" bson:"description"`
    Time        time.Time     `json:"time" bson:"time"`
    ImageURL    string        `json:"image_url" bson:"image_url"` 
}