package city

import (
    "google_io_demo/model"
)

type Repository interface {
    FindAll() ([]model.City, error)
    FindByID(cityID string) (*model.City, error)
    Save(city *model.City) error
    Update(city *model.City) error
    Delete(cityID string) error
}