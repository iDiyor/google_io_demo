package city

import (
    "google_io_demo/venue"
)

type Context struct {
    city Repository
    venue venue.Repository
}

func NewContext(city Repository, venue venue.Repository) *Context {
    return &Context{
        city: city,
        venue: venue,
    }
}