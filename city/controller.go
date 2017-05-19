package city

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
    return &Controller {
        ctx: ctx,
    }
}

func (c *Controller) handleGETCities(w http.ResponseWriter, r *http.Request) {

    cities, err := c.ctx.city.FindAll()
    if err != nil {
        panic(err)
    }

    response, err := json.Marshal(cities)
    if err != nil {
        panic(err)
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(response)
}

func (c *Controller) handlePOSTCity(w http.ResponseWriter, r *http.Request) {
    city := &model.City{}

    if err := json.NewDecoder(r.Body).Decode(city); err != nil {
        panic(err)
    }

    if err := c.ctx.city.Save(city); err != nil {
        panic(err)
    }

    response, err := json.Marshal(city)
    if err != nil {
        panic(err)
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(response)
}

func (c *Controller) handleGETCityByID(w http.ResponseWriter, r *http.Request) {

    vars := mux.Vars(r)
    cityID := vars["city_id"]

    city, err := c.ctx.city.FindByID(cityID)
    if err != nil {
        panic(err)
    }

    response, err := json.Marshal(city)
    if err != nil {
        panic(err)
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(response)
}

func (c *Controller) handleGETVenues(w http.ResponseWriter, r *http.Request) {

    vars := mux.Vars(r)
    cityID := vars["city_id"]

    venues, err := c.ctx.venue.FindAllByCityID(cityID)
    if err != nil {
        panic(err)
    }

    response, err := json.Marshal(venues)
    if err != nil {
        panic(err)
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(response)
}

func (c *Controller) handlePOSTVenue(w http.ResponseWriter, r *http.Request) {

    venue := &model.Venue{}

    if err := json.NewDecoder(r.Body).Decode(venue); err != nil {
        panic(err)
    }

    if err := c.ctx.venue.Save(venue); err != nil {
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