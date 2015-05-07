package usermodel

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type User struct {
	Id        bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Email     string        `json:"email,omitempty" bson:"email,omitempty"`
	Password  string        `json:"password,omitempty" bson:"password,omitempty"`
	CreatedAt time.Time     `json:"createdAt,omitempty" bson:"createdAt"`
}

func (u User) IsValid() bool {
	if u.Email == "" || len(u.Email) < 5 {
		return false
	}

	return true
}
