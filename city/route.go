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
    citiesRoute.HandleFunc("/", ctrl.handleGETCities)
    // GET - /cities/{city_id}
    citiesRoute.HandleFunc("/{city_id}", ctrl.handleGETCityByID)
    // GET - /cities/{city_id}/venues 
    citiesRoute.HandleFunc("/{city_id}/venues", ctrl.handleGETVenues)
}