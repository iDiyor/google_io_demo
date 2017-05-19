package event

import (
    "google_io_demo/model"
)

type Repository interface {
    FindAllByVenueID(venueID string) ([]model.Event, error)
    Save(event *model.Event) error
}