package main

import (
    "flag"
    "log"
    "net/http"  

    "google_io_demo/city"
    "google_io_demo/venue"
    "google_io_demo/deal"
    "google_io_demo/event"
    "google_io_demo/config"
    "google_io_demo/mongo"

    "gopkg.in/mgo.v2"
    "github.com/gorilla/mux"
)

var (
    port = flag.String("p", ":4000", "Server Port")
    env = flag.String("env", "dev", "Config environment")
)

func main() {
    flag.Parse()

    // Load configs
    config.Load(*env)

    // MongoDB session
    dbName := config.Current().MongoDBName
    mongoAddress := config.Current().MongoAddress
    session, err := mgo.Dial(mongoAddress)
    if err != nil {
        panic(err)
    }
    log.Printf("MongoDB Address: %s", mongoAddress)
    defer session.Close()

    session.SetMode(mgo.Monotonic, true)

    // Repositories
    var (
        cityRepo city.Repository
        venueRepo venue.Repository
        dealRepo deal.Repository
        eventRepo event.Repository
    )

    cityRepo = mongo.NewCityRepository(dbName, session)
    venueRepo, err = mongo.NewVenueRepository(dbName, session)
    checkError(err)
    dealRepo, err = mongo.NewDealRepository(dbName, session)
    checkError(err)
    eventRepo, err = mongo.NewEventRepository(dbName, session)
    checkError(err)


    // router
    router := initRoutes(cityRepo, venueRepo, dealRepo, eventRepo)
    http.Handle("/", logger(router))

    // Server
    log.Printf("Listening on port: %s", *port)
    
    err = http.ListenAndServe(*port, nil) 
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

func initRoutes(
        cityRepo city.Repository,
        venueRepo venue.Repository,
        dealRepo deal.Repository,
        eventRepo event.Repository) *mux.Router {

            router := mux.NewRouter()

            // /cities/ routes
            city.InitRoutes(router, cityRepo, venueRepo)
            
            // /venues/ routes
            venue.InitRoutes(router, venueRepo, dealRepo, eventRepo)

            return router
}

// logger middleware
func logger(handler http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
	    handler.ServeHTTP(w, r)
    })
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}