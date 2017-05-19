package mongo

import (
    "google_io_demo/venue"
    "google_io_demo/model"

    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)

type venueRepository struct {
    dbName string
    session *mgo.Session
}

func NewVenueRepository(dbName string, session *mgo.Session) (venue.Repository, error) {
    repo := &venueRepository {
        dbName: dbName,
        session: session,
    }

    venueByCityIndex := mgo.Index {
        Key:        []string {"city_id"},
        Unique:     false, // same city can have many venues - not unique
        DropDups:   false,
        Background: true,
        Sparse:     true,
        Name:       "venue_by_city",
    }

    sess := session.Copy()
    defer sess.Close()

    venueCollection := sess.DB(dbName).C("venues")

    if err := venueCollection.EnsureIndex(venueByCityIndex); err != nil {
        return nil, err
    }

    return repo, nil
}

func (r *venueRepository) FindAllByCityID(cityID string) ([]model.Venue, error) {
    sess := r.session.Copy()
    defer sess.Close()

    c := sess.DB(r.dbName).C("venues")

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

    c := sess.DB(r.dbName).C("venues")

    venue := &model.Venue{}

    err := c.Find(bson.M{"_id": bson.ObjectIdHex(venueID)}).One(venue)
    if err != nil {
        return nil, err
    }
    
    return venue, nil
}

func (r *venueRepository) Save(venue *model.Venue) error {
    sess := r.session.Copy()
    defer sess.Close()

    c := sess.DB(r.dbName).C("venues")

    err := c.Insert(venue)
    if err != nil {
        return err
    }

    return nil
}
func (r *venueRepository) Update(venue *model.Venue) error {
    sess := r.session.Copy()
    defer sess.Close()

    c := sess.DB(r.dbName).C("venues")

    _, err := c.Upsert(bson.M{"_id": venue.ID}, bson.M{"$set": venue})
    if err != nil {
        return err
    }

    return nil
}
func (r *venueRepository) Delete(venueID string) error {
    sess := r.session.Copy()
    defer sess.Close()

    c := sess.DB(r.dbName).C("venues")

    if err := c.Remove(bson.M{"_id": bson.ObjectIdHex(venueID)}); err != nil {
        panic(err)
    }

    return nil
}