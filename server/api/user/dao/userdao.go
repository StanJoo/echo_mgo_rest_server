package userdao

import (
	user "echo_mgo_rest_server/server/api/user/model"
	dbconfig "echo_mgo_rest_server/server/config"
	"errors"
	"gopkg.in/mgo.v2/bson"
)

const col string = "users"

func All() (user.Users, error) {
	db := dbconfig.DB{}
	us := user.Users{}

	s, err := db.DoDial()

	if err != nil {
		return us, errors.New("There was an error trying to connect with the DB.")
	}

	defer s.Close()

	c := s.DB(db.Name()).C(col)

	err = c.Find(bson.M{}).All(&us)

	if err != nil {
		return us, errors.New("There was an error trying to find the todos.")
	}

	return us, err
}

func GetById(id string) (user.User, error) {
	db := dbconfig.DB{}
	u := user.User{}

	s, err := db.DoDial()

	if err != nil {
		return u, errors.New("There was an error trying to connect with the DB.")
	}

	defer s.Close()

	c := s.DB(db.Name()).C(col)

	err = c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&u)

	if err != nil {
		return u, errors.New("There was an error trying to find the todos.")
	}

	return u, err
}

func NewUser(u user.User) (user.User, error) {
	db := dbconfig.DB{}

	//u.ID = bson.NewObjectId()
	u.ID = bson.NewObjectId()

	s, err := db.DoDial()

	if err != nil {
		return u, errors.New("There was an error trying to connect to the DB.")
	}

	defer s.Close()

	c := s.DB(db.Name()).C(col)

	err = c.Insert(u)

	if err != nil {
		return u, errors.New("There was an error trying to insert the doc to the DB.")
	}

	return u, err
}

func DeleteUser(id string) error {
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
		return errors.New("There was an error trying to remove the todo.")
	}

	return err
}
