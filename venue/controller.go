package venue

import (
    "net/http"
    "encoding/json"

    "google_io_demo/model"

    "github.com/gorilla/mux"
)

type Controller struct {
    ctx *Context
}

func NewController(ctx *Context) *Controller {
    return &Controller{
        ctx: ctx,
    }
}

func (c *Controller) handleGETVenueByID(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    venueID := vars["venue_id"]


    venue, err := c.ctx.venue.FindByID(venueID)
    if err != nil {
        panic(err)
    }

    response, err := json.Marshal(venue)
    if err != nil {
        panic(err)
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(response)
}

func (c *Controller) handlePUTVenue(w http.ResponseWriter, r *http.Request) {

    venue := &model.Venue{}

    if err := json.NewDecoder(r.Body).Decode(venue); err != nil {
        panic(err)
    }

    if err := c.ctx.venue.Update(venue); err != nil {
        panic(err)
    }

    response, err := json.Marshal(venue)
    if err != nil {
        panic(err)
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(response)
}

func (c *Controller) handleDeleteVenue(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    venueID := vars["venue_id"]


    if err := c.ctx.venue.Delete(venueID); err != nil {
        panic(err)
    }

    message := map[string]string {
        "message": "Success",
    }

    response, err := json.Marshal(message)
    if err != nil {
        panic(err)
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(response)
}

func (c *Controller) handlePOSTDeal(w http.ResponseWriter, r *http.Request) {

    deal := &model.Deal{}

    if err := json.NewDecoder(r.Body).Decode(deal); err != nil {
        panic(err)
    }

    if err := c.ctx.deal.Save(deal); err != nil {
        panic(err)
    }

    response, err := json.Marshal(deal)
    if err != nil {
        panic(err)
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(response)
}


func (c *Controller) handleGETDeal(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    venueID := vars["venue_id"]

    deals, err := c.ctx.deal.FindAllByVenueID(venueID)
    if err != nil {
        panic(err)
    }

    response, err := json.Marshal(deals)
    if err != nil {
        panic(err)
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(response)
}

func (c *Controller) handlePOSTEvent(w http.ResponseWriter, r *http.Request) {

    event := &model.Event{}

    if err := json.NewDecoder(r.Body).Decode(event); err != nil {
        panic(err)
    }

    if err := c.ctx.event.Save(event); err != nil {
        panic(err)
    }

    response, err := json.Marshal(event)
    if err != nil {
        panic(err)
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(response)
}

func (c *Controller) handleGETEvent(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    venueID := vars["venue_id"]

    events, err := c.ctx.event.FindAllByVenueID(venueID)
    if err != nil {
        panic(err)
    }

    response, err := json.Marshal(events)
    if err != nil {
        panic(err)
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(response)
}