package mongo

import (

    "google_io_demo/event"
    "google_io_demo/model"

    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)

type venueRepository struct {
    dbName string
    session *mgo.Session
}

func NewVenueRepository(dbName string, session *mgo.Session) (event.Repository, error) {
    repo := &eventRepository {
        dbName: dbName,
        session: session,
    }

    return repo, nil
}

func (r *venueRepository) FindAllByCityID(cityID string) ([]model.Venue, error) {
    sess := r.session.Copy()
    defer sess.Close()

    c := sess.DB(r.dbName).C("venue")

    venues := make([]model.Venue, 0)

    err := c.Find(bson.M{"city_id": cityID}).All(&venues)
    if err != nil {
        return nil, err
    }
    
    return venues, nil
}

func (r *venueRepository) FindByID(venueID string) (*model.Venue, error) {
    sess := r.session.Copy()
    defer sess.Close()

    c := sess.DB(r.dbName).C("venue")

    venue := &model.Venue{}

    err := c.Find(bson.M{"_id": venueID}).One(venue)
    if err != nil {
        return nil, err
    }
    
    return venue, nil
}