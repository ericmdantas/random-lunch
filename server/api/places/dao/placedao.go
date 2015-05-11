package placedao

import (
	"errors"
	places "github.com/ericmdantas/random-lunch/server/api/places/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	URL_DB  string = "localhost"
	DB_NAME        = "random-lunch"
	COLL           = "places"
)

func All() ([]places.Place, error) {
	var ps []places.Place

	session, err := mgo.Dial(URL_DB)
	defer session.Close()

	if err != nil {
		return ps, errors.New("Error while trying to open connection with DB.")
	}

	err = session.DB(DB_NAME).C(COLL).Find(bson.M{}).All(&ps)

	if err != nil {
		return ps, errors.New("Error while trying to retrieve all places.")
	}

	return ps, nil
}

func CreateNew(p *places.Place) error {

	if !p.IsValid() {
		return errors.New("Invalid place.")
	}

	session, err := mgo.Dial(URL_DB)
	defer session.Close()

	if err != nil {
		return errors.New("Error while opening connection with DB.")
	}

	err = session.DB(DB_NAME).C(COLL).Insert(p)

	if err != nil {
		return errors.New("Error while inserting place in DB.")
	}

	return nil
}

func Update(pl *places.Place) error {
	if pl.IsNew() {
		return errors.New("Error while trying to update the place - Place has no Id.")
	}

	session, err := mgo.Dial(URL_DB)
	defer session.Close()

	if err != nil {
		return errors.New("Error while trying to connect to the DB.")
	}

	err = session.DB(DB_NAME).C(COLL).Update(bson.M{"_id": pl.Id}, pl)

	if err != nil {
		return errors.New("Error while trying to update the place.")
	}

	return nil
}
