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

    eventByVenueIndex := mgo.Index {
        Key:        []string {"venue_id"},
        Unique:     false, // same venue can have many events - not unique
        DropDups:   false,
        Background: true,
        Sparse:     true,
        Name:       "event_by_venue",
    }

    sess := session.Copy()
    defer sess.Close()

    eventCollection := sess.DB(dbName).C("events")

    if err := eventCollection.EnsureIndex(eventByVenueIndex); err != nil {
        return nil, err
    }

    return repo, nil
}

func (r *eventRepository) FindAllByVenueID(venueID string) ([]model.Event, error) {
    sess := r.session.Copy()
    defer sess.Close()

    c := sess.DB(r.dbName).C("events")

    events := make([]model.Event, 0)

    err := c.Find(bson.M{"venue_id": venueID}).All(&events)
    if err != nil {
        return nil, err
    }
    
    return events, nil
}

func (r *eventRepository) Save(event *model.Event) error {
    sess := r.session.Copy()
    defer sess.Close()

    c := sess.DB(r.dbName).C("events")

     err := c.Insert(event)
     if err != nil {
         return err
     }

     return nil
}