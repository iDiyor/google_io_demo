package venue

import (
    "google_io_demo/deal"
    "google_io_demo/event"
)

type Context struct {
    venue Repository
    deal deal.Repository
    event event.Repository
}

func NewContext(venue Repository, deal deal.Repository, event event.Repository) *Context {
    return &Context {
        venue: venue,
        deal: deal,
        event: event,
    }
}
