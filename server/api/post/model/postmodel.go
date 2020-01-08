package postmodel

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Post struct {
	ID       bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Title    string        `json:"title" bson:"title"`
	Article  string        `json:"article,omitempty" bson:"article"`
	CreateAt time.Time     `json:"createdAt,omitempty" bson:"createdAt"`
}

type Posts []Post
