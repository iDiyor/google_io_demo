package city

import (

    "google_io_demo/venue"

    "github.com/gorilla/mux"
)

func InitRoutes(router *mux.Router, city Repository, venue venue.Repository) {
    ctx := NewContext(city, venue)
    ctrl := NewController(ctx)

    citiesRoute := router.PathPrefix("/cities").Subrouter()
    
    // GET - /cities/
    citiesRoute.HandleFunc("/", ctrl.handleGETCities).Methods("GET")
    // POST - /cities/
    citiesRoute.HandleFunc("/", ctrl.handlePOSTCity).Methods("POST")
    // GET - /cities/{city_id}
    citiesRoute.HandleFunc("/{city_id}", ctrl.handleGETCityByID).Methods("GET")
    // GET - /cities/{city_id}/venues 
    citiesRoute.HandleFunc("/{city_id}/venues", ctrl.handleGETVenues).Methods("GET")
    // POST - /cities/{city_id}/venues 
    citiesRoute.HandleFunc("/{city_id}/venues", ctrl.handlePOSTVenue).Methods("POST")
}