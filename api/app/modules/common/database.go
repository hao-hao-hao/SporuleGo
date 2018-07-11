package common

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//Resource is the interface for dbs
type Resource interface {
	Create(table string, item interface{}) error
	Get(table string, object, query interface{}, extraQuery func(*mgo.Query) *mgo.Query) error
	GetAll(table string, objects, query interface{}, extraQuery func(*mgo.Query) *mgo.Query) error
	Update(table string, query, updatedItem interface{}, updateAll bool) error
	Delete(table string, query interface{}, RemoveAll bool) error
}

//MongoDB Type provides the basic connection string
type MongoDB struct {
	session *mgo.Session
	//Session provides a copied session for operations, remember to close it by using defer session.Close()
	Session func() *mgo.Session
}

//Resources is the db
var Resources Resource

//NewMongoDB Creates an session object
func NewMongoDB(host, database, username, password string, dropDB bool) (*MongoDB, error) {
	db := &MongoDB{}
	var err error
	//create a new session
	db.session, err = mgo.DialWithInfo(
		&mgo.DialInfo{
			Username: username,
			Password: password,
			Database: database,
			Addrs:    []string{host},
		},
	)
	if err != nil {
		return nil, err
	}
	//set up the Session fucntion to return a copy of the db session
	db.Session = func() *mgo.Session { return db.session.Copy() }
	if dropDB {
		//drop the database for testing
		db.dropDatabase()
	}
	return db, nil
}

//collection provides a copied session as well as a collection, this is a helper function for CRUD
func (db *MongoDB) collection(collection string) (*mgo.Session, *mgo.Collection) {
	s := db.Session()
	c := s.DB(Config.Database).C(collection)
	return s, c
}

//Create provides Insert Operation for Database
func (db *MongoDB) Create(collection string, item interface{}) error {
	s, c := db.collection(collection)
	defer s.Close()
	err := c.Insert(item)
	return err
}

//Get gets single item that matches query, for example bson.M{"_id": id}
func (db *MongoDB) Get(table string, object, query interface{}, extraQuery func(*mgo.Query) *mgo.Query) error {
	s, c := db.collection(table)
	defer s.Close()
	q := c.Find(query)
	if extraQuery != nil {
		q = extraQuery(q)
	}
	err := q.One(object)
	return err
}

//GetAll gets all item that matches query, for example bson.M{"Name": "Hello"}
func (db *MongoDB) GetAll(table string, objects, query interface{}, extraQuery func(*mgo.Query) *mgo.Query) error {
	s, c := db.collection(table)
	defer s.Close()
	q := c.Find(query)
	if extraQuery != nil {
		q = extraQuery(q)
	}
	err := q.All(objects)
	return err
}

//Update provides Update Operation for Database
func (db *MongoDB) Update(collection string, query, updatedItem interface{}, UpdateAll bool) error {
	s, c := db.collection(collection)
	defer s.Close()
	var err error
	if UpdateAll {
		_, err = c.UpdateAll(query, bson.M{"$set": updatedItem})
	} else {
		err = c.Update(query, updatedItem)
	}
	return err
}

//Delete provides Delete Operation for Database
func (db *MongoDB) Delete(collection string, query interface{}, RemoveAll bool) error {
	s, c := db.collection(collection)
	defer s.Close()
	var err error
	if RemoveAll {
		_, err = c.RemoveAll(query)
		return err
	}
	err = c.Remove(query)
	return err
}

//DropDatabase drop the database
func (db *MongoDB) dropDatabase() error {
	return db.session.DB(Config.Database).DropDatabase()
}
