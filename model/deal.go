package model

type Deal struct {
    ID          string `json:"id" bson:"_id,omitempty"`
    VenueID     string `json:"venue_id"`
    Name        string `json:"name"`
    Description string `json:"description"`
    ImageURL    string `json:"image_url"` 
}