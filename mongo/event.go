package mongo

import (

    "google_io_demo/event"
    "google_io_demo/model"

    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)

type eventRepository struct {
    dbName string
    session *mgo.Session
}

func NewEventRepository(dbName string, session *mgo.Session) (event.Repository, error) {
    repo := &eventRepository {
        dbName: dbName,
        session: session,
    }

    return repo, nil
}

func (r *eventRepository) FindAllByVenueID(venueID string) ([]model.Event, error) {
    sess := r.session.Copy()
    defer sess.Close()

    c := sess.DB(r.dbName).C("event")

    events := make([]model.Event, 0)

    err := c.Find(bson.M{"venue_id": venueID}).All(&events)
    if err != nil {
        return nil, err
    }
    
    return events, nil
}