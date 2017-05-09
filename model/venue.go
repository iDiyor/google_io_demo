package model

type Venue struct {
	ID          string `json:"id" bson:"_id,omitempty"`
	CityID      string `json:"city_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
    Address     string `json:"address"`
    ImageURL    string `json:"image_url"`
}
