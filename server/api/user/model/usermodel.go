package usermodel

import (
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID       bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Username string        `json:"username" bson:"username"`
	Password string        `json:"password,omitempty" bson:"password"`
	Token    string        `json:"token,omitempty" bson:"-"`
}

func (u User) IsValid() bool {
	if l := len(u.Password); l >= 8 {
		return true
	}

	return false
}

type Users []User
