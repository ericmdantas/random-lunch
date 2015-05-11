package placemodel

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Place struct {
	Id             bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Name           string        `json:"name,omitempty" bson:"name,omitempty"`
	Address        string        `json:"address,omitempty" bson:"address,omitempty"`
	Image          string        `json:"image,omitempty" bson:"image,omitempty"`
	LastTimeChosen time.Time     `json:"lastTimeChosen,omitempty" bson:"lastTimeChosen,omitempty"`
	CreatedAt      time.Time     `json:"createdAt,omitempty" bson:"createdAt"`
}

func (p Place) IsValid() bool {

	if p.Name == "" || p.Address == "" {
		return false
	}

	return true
}

func (p Place) IsNew() bool {
	if p.Id != "" {
		return false
	}

	return true
}
