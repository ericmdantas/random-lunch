package userdao

import (
	"errors"
	user "github.com/ericmdantas/random-lunch/server/api/users/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	URL_DB  string = "localhost"
	DB_NAME string = "random-lunch"
	COLL    string = "users"
)

func All(us *[]user.User) error {
	session, err := mgo.Dial(URL_DB)
	defer session.Close()

	if err != nil {
		return errors.New("Error while trying to open the connection with the DB.")
	}

	err = session.DB(DB_NAME).C(COLL).Find(bson.M{}).All(us)

	if err != nil {
		return errors.New("Error while trying to retrieve all users.")
	}

	return nil
}

func CreateNew(us *user.User) error {

	if !us.IsValid() {
		return errors.New("Invalid user.")
	}

	session, err := mgo.Dial(URL_DB)
	defer session.Close()

	if err != nil {
		panic(err)
	}

	us.Id = bson.NewObjectId()

	err = session.DB(DB_NAME).C(COLL).Insert(us)

	if err != nil {
		return errors.New("Error while trying to insert user.")
	}

	return nil
}
