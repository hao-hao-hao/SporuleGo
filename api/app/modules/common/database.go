package common

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//Resource is the interface for dbs, but we are not using it for now
type Resource interface {
	Create(table string, item interface{}) error
	Get(table string, object, query interface{}, extraQuery func(*mgo.Query) *mgo.Query) error
	GetAll(table string, objects, query interface{}, extraQuery func(*mgo.Query) *mgo.Query) error
	Update(table string, query, updatedItem interface{}, updateAll bool) error
	Delete(table string, query interface{}, RemoveAll bool) error
	AggGet(table string, object interface{}, queries ...bson.M) error
	AggGetAll(table string, objects interface{}, queries ...bson.M) error
}

//MongoDB Type is simply a holder
type MongoDB struct {
	//This original session is not open to public
	session *mgo.Session
	//Session provides a copied session for operations, remember to close it by using defer session.Close()
	Session func() *mgo.Session
}

//Resources is the db
var Resources *MongoDB

//NewMongoDB initiates the db session
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
		//This is purely for testing purpose, it will drop the database if it is true.
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
	//get the session and collection
	s, c := db.collection(collection)
	//close the session after usage
	defer s.Close()
	//run the Insert function
	err := c.Insert(item)
	return err
}

//AggGet is the aggregate pipe function for mongo db, it takes an bson.M arrary query and assign one item to object.
func (db *MongoDB) AggGet(table string, object interface{}, queries ...bson.M) error {
	s, c := db.collection(table)
	defer s.Close()
	err := c.Pipe(queries).One(object)
	return err
}

//AggGetAll is the aggregate pipe function for mongo db, it takes an bson.M arrary query and assign an arrary to the objects
func (db *MongoDB) AggGetAll(table string, objects interface{}, queries ...bson.M) error {
	s, c := db.collection(table)
	defer s.Close()
	err := c.Pipe(queries).All(objects)
	return err
}

//Get takes in the table name, pointer to the single obejct,
//mongodb query and a hook function to implement the extra query such as select,slice..
//the result will be return to the object pointer
func (db *MongoDB) Get(table string, object, query interface{}, extraQuery func(*mgo.Query) *mgo.Query) error {
	s, c := db.collection(table)
	defer s.Close()
	q := c.Find(query)
	if extraQuery != nil {
		//run queries in the hook if the hook is not empty
		q = extraQuery(q)
	}
	err := q.One(object)
	return err
}

//GetAll takes in the table name, pointer to the obejct array,
//mongodb query and a hook function to implement the extra query such as select,slice..
//the result will be return to the object array pointer
func (db *MongoDB) GetAll(table string, objects, query interface{}, extraQuery func(*mgo.Query) *mgo.Query) error {
	s, c := db.collection(table)
	defer s.Close()
	q := c.Find(query)
	if extraQuery != nil {
		//run queries in the hook if the hook is not empty
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
