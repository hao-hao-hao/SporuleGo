package common

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type testStruct struct {
	Name string `bson:"name"`
}

func TestNewMongoDB(t *testing.T) {
	//Load Test settings
	Config.LoadConfiguration("../../../config/test.json")
	convey.Convey("TestNewMongoDB should return a non nil db with session, and non nil Session() function", t, func() {
		db, _ := NewMongoDB(Config.Host, Config.Database, Config.Username, Config.Password, Config.DropDB)
		convey.So(db.Session, convey.ShouldNotBeNil)
		convey.So(db.Session, convey.ShouldNotBeNil)
	})
}
func TestCreate(t *testing.T) {
	//Load Test settings
	Config.LoadConfiguration("../../../config/test.json")
	convey.Convey("Create should create an item without error", t, func() {
		db := &MongoDB{}
		var err error
		db.session, err = mgo.DialWithInfo(
			&mgo.DialInfo{
				Username: Config.Username,
				Password: Config.Password,
				Database: Config.Database,
				Addrs:    []string{Config.Host},
			},
		)
		convey.So(err, convey.ShouldBeNil)
		if Config.DropDB {
			//drop the database for testing
			db.session.DB(Config.Database).DropDatabase()
		}
		db.Session = func() *mgo.Session { return db.session.Copy() }
		collection := "test"
		err = db.Create(collection, &testStruct{Name: "hello"})
		convey.So(err, convey.ShouldBeNil)
	})

}

func TestGet(t *testing.T) {
	//Load Test settings
	Config.LoadConfiguration("../../../config/test.json")
	convey.Convey("It should return an item that matches the criteria name hello", t, func() {
		db := &MongoDB{}
		var err error
		db.session, err = mgo.DialWithInfo(
			&mgo.DialInfo{
				Username: Config.Username,
				Password: Config.Password,
				Database: Config.Database,
				Addrs:    []string{Config.Host},
			},
		)
		convey.So(err, convey.ShouldBeNil)
		if Config.DropDB {
			//drop the database for testing
			db.session.DB(Config.Database).DropDatabase()
		}
		db.Session = func() *mgo.Session { return db.session.Copy() }
		collection := "test"
		err = db.Session().DB(Config.Database).C(collection).Insert(&testStruct{Name: "Hello"})
		err = db.Session().DB(Config.Database).C(collection).Insert(&testStruct{Name: "ABC"})

		//get by name = hello
		var tempTest testStruct
		err = db.Get(collection, &tempTest, bson.M{"name": "Hello"})
		convey.So(tempTest.Name, convey.ShouldEqual, "Hello")

	})
}

func TestGetAll(t *testing.T) {
	//Load Test settings
	Config.LoadConfiguration("../../../config/test.json")
	convey.Convey("It should return the 2 items which names are Hello", t, func() {
		db := &MongoDB{}
		var err error
		db.session, err = mgo.DialWithInfo(
			&mgo.DialInfo{
				Username: Config.Username,
				Password: Config.Password,
				Database: Config.Database,
				Addrs:    []string{Config.Host},
			},
		)
		convey.So(err, convey.ShouldBeNil)
		if Config.DropDB {
			//drop the database for testing
			db.session.DB(Config.Database).DropDatabase()
		}
		db.Session = func() *mgo.Session { return db.session.Copy() }
		collection := "test"
		err = db.Session().DB(Config.Database).C(collection).Insert(&testStruct{Name: "Hello"})
		err = db.Session().DB(Config.Database).C(collection).Insert(&testStruct{Name: "Hello"})
		err = db.Session().DB(Config.Database).C(collection).Insert(&testStruct{Name: "ABC"})

		//get by name = hello
		var tempTests []testStruct
		err = db.GetAll(collection, &tempTests, bson.M{"name": "Hello"})
		convey.So(len(tempTests), convey.ShouldEqual, 2)

	})
}

func TestUpdate(t *testing.T) {
	//Load Test settings
	Config.LoadConfiguration("../../../config/test.json")
	convey.Convey("Testing update and update all", t, func() {
		db := &MongoDB{}
		var err error
		db.session, err = mgo.DialWithInfo(
			&mgo.DialInfo{
				Username: Config.Username,
				Password: Config.Password,
				Database: Config.Database,
				Addrs:    []string{Config.Host},
			},
		)
		convey.So(err, convey.ShouldBeNil)
		if Config.DropDB {
			//drop the database for testing
			db.session.DB(Config.Database).DropDatabase()
		}
		db.Session = func() *mgo.Session { return db.session.Copy() }
		collection := "test"
		c := db.Session().DB(Config.Database).C(collection)
		err = c.Insert(&testStruct{Name: "ABC"})
		err = c.Insert(&testStruct{Name: "ABC"})
		err = c.Insert(&testStruct{Name: "ABC"})
		err = c.Insert(&testStruct{Name: "BBC"})
		var temps []testStruct
		convey.Convey("Update One should only update one item, so the length should be 1, it also check if the returned item is updated", func() {
			err = db.Update(collection, bson.M{"name": "ABC"}, &testStruct{Name: "Hello"}, false)
			err = c.Find(bson.M{"name": "Hello"}).All(&temps)
			convey.So(len(temps), convey.ShouldEqual, 1)
			convey.So(temps[0].Name, convey.ShouldEqual, "Hello")
		})

		convey.Convey("Update All should update all items, so the length should be 3", func() {
			err = db.Update(collection, bson.M{"name": "ABC"}, &testStruct{Name: "Hello"}, true)
			err = c.Find(bson.M{"name": "Hello"}).All(&temps)
			convey.So(len(temps), convey.ShouldEqual, 3)
			convey.So(temps[0].Name, convey.ShouldEqual, "Hello")
		})

	})
}

func TestDelete(t *testing.T) {
	//Load Test settings
	Config.LoadConfiguration("../../../config/test.json")
	convey.Convey("Testing the delete and delete all", t, func() {
		db := &MongoDB{}
		var err error
		db.session, err = mgo.DialWithInfo(
			&mgo.DialInfo{
				Username: Config.Username,
				Password: Config.Password,
				Database: Config.Database,
				Addrs:    []string{Config.Host},
			},
		)
		convey.So(err, convey.ShouldBeNil)
		if Config.DropDB {
			//drop the database for testing
			db.session.DB(Config.Database).DropDatabase()
		}
		db.Session = func() *mgo.Session { return db.session.Copy() }
		collection := "test"
		c := db.Session().DB(Config.Database).C(collection)
		err = c.Insert(&testStruct{Name: "ABC"})
		err = c.Insert(&testStruct{Name: "ABC"})
		err = c.Insert(&testStruct{Name: "ABC"})
		err = c.Insert(&testStruct{Name: "BBC"})
		var temps []testStruct

		convey.Convey("It should return 3 item which the name is ABC", func() {
			err = db.Delete(collection, bson.M{"name": "ABC"}, false)
			err = c.Find(nil).All(&temps)
			convey.So(len(temps), convey.ShouldEqual, 3)
		})

		convey.Convey("It should return 1 item which the name is BBC", func() {
			err = db.Delete(collection, bson.M{"name": "ABC"}, true)
			err = c.Find(nil).All(&temps)
			convey.So(len(temps), convey.ShouldEqual, 1)
			convey.So(temps[0].Name, convey.ShouldEqual, "BBC")
		})

	})
}
