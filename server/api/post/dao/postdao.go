package postdao

import (
	post "echo_mgo_rest_server/server/api/post/model"
	dbconfig "echo_mgo_rest_server/server/config"
	"errors"
	"gopkg.in/mgo.v2/bson"
	"time"
)

const col string = "posts"

func All() (post.Posts, error) {
	db := dbconfig.DB{}
	ps := post.Posts{}

	s, err := db.DoDial()

	if err != nil {
		return ps, errors.New("There was an error trying to connect with the DB.")
	}

	defer s.Close()

	c := s.DB(db.Name()).C(col)

	err = c.Find(bson.M{}).All(&ps)

	if err != nil {
		return ps, errors.New("There was an error trying to find the posts.")
	}

	return ps, err
}

func GetByTitle(title string) (post.Post, error) {
	db := dbconfig.DB{}
	p := post.Post{}

	s, err := db.DoDial()

	if err != nil {
		return p, errors.New("There was an error trying to connect with the DB.")
	}

	defer s.Close()

	c := s.DB(db.Name()).C(col)

	err = c.Find(bson.M{"title": bson.ObjectIdHex(title)}).One(&p)

	if err != nil {
		return p, errors.New("There was an error trying to find the post.")
	}

	return p, err
}

func NewPost(p post.Post) (post.Post, error) {
	db := dbconfig.DB{}

	p.ID = bson.NewObjectId()
	p.CreateAt = time.Now()

	s, err := db.DoDial()

	if err != nil {
		return p, errors.New("There was an error trying to connect to the DB.")
	}

	defer s.Close()

	c := s.DB(db.Name()).C(col)

	err = c.Insert(p)

	if err != nil {
		return p, errors.New("There was an error trying to insert the doc to the DB.")
	}

	return p, err
}

func DeletePost(id string) error {
	db := dbconfig.DB{}

	s, err := db.DoDial()

	if err != nil {
		return errors.New("There was an error trying to connect with the DB.")
	}

	idoi := bson.ObjectIdHex(id)

	defer s.Close()

	c := s.DB(db.Name()).C(col)

	err = c.RemoveId(idoi)

	if err != nil {
		return errors.New("There was an error trying to remove the post.")
	}
	return err
}
