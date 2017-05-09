package model

import "time"

type Event struct {
    ID          string    `json:"id" bson:"_id,omitempty"`
    VenueID     string    `json:"venue_id"`
    Name        string    `json:"name"`
    Description string    `json:"description"`
    Time        time.Time `json:"time"`
    ImageURL    string    `json:"image_url"` 
}