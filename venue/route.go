package venue

import (

    "google_io_demo/deal"
    "google_io_demo/event"

    "github.com/gorilla/mux"
)

func InitRoutes(router *mux.Router, venue Repository, deal deal.Repository, event event.Repository) {

    ctx := NewContext(venue, deal, event)
    ctrl := NewController(ctx)

    venuesRoute := router.PathPrefix("/venues").Subrouter()
    
    // GET - /venues/{venue_id}
    venuesRoute.HandleFunc("/{venue_id}", ctrl.handleGETVenueByID).Methods("GET")

    // PUT - /venues/{venue_id}
    venuesRoute.HandleFunc("/{venue_id}", ctrl.handlePUTVenue).Methods("PUT")

    // DELETE - /venues/{venue_id}
    venuesRoute.HandleFunc("/{venue_id}", ctrl.handleDeleteVenue).Methods("DELETE")

    // POST - /venues/{venue_id}/deals
    venuesRoute.HandleFunc("/{venue_id}/deals", ctrl.handlePOSTDeal).Methods("POST")

    // GET - /venues/{venue_id}/deals
    venuesRoute.HandleFunc("/{venue_id}/deals", ctrl.handleGETDeal).Methods("GET")

    // POST - /venues/{venue_id}/events
    venuesRoute.HandleFunc("/{venue_id}/events", ctrl.handlePOSTEvent).Methods("POST")

    // GET - /venues/{venue_id}/events
    venuesRoute.HandleFunc("/{venue_id}/events", ctrl.handleGETEvent).Methods("GET")
}