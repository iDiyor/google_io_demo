package venue

import (
    "google_io_demo/model"
)

type Repository interface {
    FindAllByCityID(cityID string) ([]model.Venue, error)
    FindByID(venueID string) (*model.Venue, error)
}