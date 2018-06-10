package models

import (
	"errors"
	"sporule/api/app/modules/common"

	"gopkg.in/mgo.v2/bson"
)

//fieldCollection is the collection name for Model Node in mongo db
const fieldCollection = "field"

//Field is for the purpose of front end rendering. Such as dropdown or textbox or input string etc....
type Field struct {
	ID        bson.ObjectId `bson:"_id"`
	Name      string        `bson:"name"`
	FieldType string        `bson:"type"`
}

//NewField is the constructor for Field
func NewField(name, fieldType string) (*Field, error) {
	if !common.CheckNil(name, fieldType) {
		return nil, errors.New(common.Enums.ErrorMessages.LackOfInfo)
	}
	field := &Field{}
	field.ID = bson.NewObjectId()
	field.Name = name
	field.FieldType = fieldType
	return field, nil
}

//Insert inserts the field to the database
func (field *Field) Insert() error {
	if !common.CheckNil(field.Name, field.FieldType) {
		return errors.New(common.Enums.ErrorMessages.LackOfInfo)
	}
	if common.Create(fieldCollection, field) != nil {
		return errors.New(common.Enums.ErrorMessages.LackOfInfo)
	}
	return nil
}

//Update updates the Field to the database
func (field *Field) Update(id bson.ObjectId) error {
	if !common.CheckNil(id) {
		id = field.ID
	} else {
		field.ID = id
	}
	err := common.Update(fieldCollection, bson.M{"_id": id}, field, false)
	return err
}

//DeleteField deletes the selected field by Id
func DeleteField(id bson.ObjectId) error {
	return common.Delete(fieldCollection, bson.M{"_id": id}, true)
}

//GetField returns a field according to the filter query
func GetField(query bson.M) (*Field, error) {
	var field Field
	s, c := common.Collection(fieldCollection)
	defer s.Close()
	err := c.Find(query).One(&field)
	return &field, err
}

//GetFields returns fields according to the filter query
func GetFields(query bson.M) (*[]Field, error) {
	var fields []Field
	s, c := common.Collection(fieldCollection)
	defer s.Close()
	err := c.Find(query).All(&fields)
	return &fields, err
}

//GetFieldByID returns field by id
func GetFieldByID(id bson.ObjectId) (*Field, error) {
	field, err := GetField(bson.M{"_id": id})
	return field, err
}
