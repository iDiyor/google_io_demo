package deal

import (
    "google_io_demo/model"
)

type Repository interface {
    FindAllByVenueID(venueID string) ([]model.Deal, error)
    Save(deal *model.Deal) error
}