package model

type City struct {
    ID      string `json:"id" bson:"_id,omitempty"`
    Name    string `json:"name"`
    Country string `json:"country"`
}