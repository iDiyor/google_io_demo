package mongo

import (

    "google_io_demo/deal"
    "google_io_demo/model"

    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)

type dealRepository struct {
    dbName string
    session *mgo.Session
}

func NewDealRepository(dbName string, session *mgo.Session) (deal.Repository, error) {
    repo := &dealRepository {
        dbName: dbName,
        session: session,
    }

    return repo, nil
}

func (r *dealRepository) FindAllByVenueID(venueID string) ([]model.Deal, error) {
    sess := r.session.Copy()
    defer sess.Close()

    c := sess.DB(r.dbName).C("deal")

    deals := make([]model.Deal, 0)

    err := c.Find(bson.M{"venue_id": venueID}).All(&deals)
    if err != nil {
        return nil, err
    }
    
    return deals, nil
}