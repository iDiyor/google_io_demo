package mongo

import (

    "google_io_demo/city"
    "google_io_demo/model"

    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)

type cityRepository struct {
    dbName string
    session *mgo.Session
}

func NewCityRepository(dbName string, session *mgo.Session) city.Repository {
    return &cityRepository {
        dbName: dbName,
        session: session,
    }
}

func (r *cityRepository) FindAll() ([]model.City, error) {
    sess := r.session.Copy()
    defer sess.Close()

    c := sess.DB(r.dbName).C("city")

    cities := make([]model.City, 0)

    err := c.Find(nil).All(&cities)
    if err != nil {
        return nil, err
    }
    
    return cities, nil
}

func (r *cityRepository) FindByID(cityID string) (*model.City, error) {
    sess := r.session.Copy()
    defer sess.Close()

    c := sess.DB(r.dbName).C("city")

    city := &model.City{}

    err := c.Find(bson.M{"_id": cityID}).One(city)
    if err != nil {
        return nil, err
    }
    
    return city, nil
}

func (r *cityRepository) Save(city *model.City) error {
    sess := r.session.Copy()
    defer sess.Close()

    c := sess.DB(r.dbName).C("city")

    err := c.Insert(city)
    if err != nil {
        return err
    }

    return nil
}

func (r *cityRepository) Update(city *model.City) error {
    return nil
}

func (r *cityRepository) Delete(cityID string) error {
    return nil
}