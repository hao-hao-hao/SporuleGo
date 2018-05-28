package common

import (
	"gopkg.in/mgo.v2"
)

//Session is the Database Session
var session *mgo.Session

//db is the database name
const db = "sporule"

//Database Type provides the basic connection string
type Database struct {
	Host, Database, Username, Password string
	IsDrop                             bool
	Timeout                            uint
}

//InitiateDatabase Creates an session object
func InitiateDatabase() {

	var database = Database{
		Host:     "localhost:27017",
		Database: db,
		Username: "sporule",
		Password: "1q2w3e4r",
		IsDrop:   false,
	}

	var err error
	//create a new session
	session, err = mgo.DialWithInfo(
		&mgo.DialInfo{
			Username: database.Username,
			Password: database.Password,
			Database: database.Database,
			Addrs:    []string{database.Host},
		},
	)
	if err != nil {
		//throw error message
		panic(err)
	}
	if database.IsDrop {
		//drop the database for testing
		session.DB(db).DropDatabase()
	}

}

//Session provides a copied session for operations, remember to close it by using defer session.Close()
func Session() *mgo.Session {
	return session.Copy()
}

//Collection provides a copied session as well as a collection, this is a helper function for CRUD
func Collection(collection string) (*mgo.Session, *mgo.Collection) {
	s := Session()
	c := s.DB(db).C(collection)
	return s, c
}

//Create provides Insert Operation for Database
func Create(collection string, item interface{}) error {
	s, c := Collection(collection)
	defer s.Close()
	err := c.Insert(item)
	return err
}

//Update provides Update Operation for Database
func Update(collection string, selector, updatedItem interface{}, isAll bool) error {
	s, c := Collection(collection)
	defer s.Close()
	var err error
	if isAll {
		_, err = c.UpdateAll(selector, updatedItem)
	} else {
		err = c.Update(selector, updatedItem)
	}
	return err
}

//Delete provides Delete Operation for Database
func Delete(collection string, selector interface{}, isAll bool) error {
	s, c := Collection(collection)
	defer s.Close()
	var err error
	if isAll {
		_, err = c.RemoveAll(selector)
	}
	err = c.Remove(selector)
	return err
}

//DropDatabase drop the database
func DropDatabase() {
	session.DB(db).DropDatabase()
}
