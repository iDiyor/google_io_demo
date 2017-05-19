package venue

import (
    "google_io_demo/model"
)

type Repository interface {
    FindAllByCityID(cityID string) ([]model.Venue, error)
    FindByID(venueID string) (*model.Venue, error)
    Save(venue *model.Venue) error
    Update(venue *model.Venue) error
    Delete(venueID string) error
}