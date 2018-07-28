package models

import (
	"errors"
	"sporule/api/app/modules/common"
	"time"

	"gopkg.in/mgo.v2/bson"
)

//fieldCollection is the collection name for Model Node in mongo db
const fieldCollection = "field"

//Field is for the purpose of front end rendering. Such as dropdown or textbox or input string etc....
type Field struct {
	ID           bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
	Name         string        `bson:"name,omitempty" json:"name,omitempty"`
	Type         string        `bson:"type,omitempty" json:"type,omitempty"`
	Preset       string        `bson:"preset,omitempty" json:"preset,omitempty"`
	CreatedDate  time.Time     `bson:"createdDate,omitempty" json:"createdDate,omitempty"`
	ModifiedDate time.Time     `bson:"modeifiedDate,omitempty" json:"modeifiedDate,omitempty"`
}

//NewField is the constructor for Field
func NewField(name, fieldType string) (*Field, error) {
	if !common.CheckNil(name, fieldType) {
		return nil, errors.New(common.Enums.ErrorMessages.LackOfInfo)
	}
	field := &Field{}
	field.ID = bson.NewObjectId()
	field.Name = name
	field.Type = fieldType
	return field, nil
}

//Insert inserts the field to the database
func (field *Field) Insert() error {
	if !common.CheckNil(field.ID, field.Name, field.Type) {
		return errors.New(common.Enums.ErrorMessages.LackOfInfo)
	}
	if field.IsExist() {
		return errors.New(common.Enums.ErrorMessages.RecordExist)
	}
	field.CreatedDate = time.Now()
	field.ModifiedDate = time.Now()
	if common.Resources.Create(fieldCollection, field) != nil {
		return errors.New(common.Enums.ErrorMessages.LackOfInfo)
	}
	return nil
}

//Update updates the Field to the database
func (field *Field) Update() error {
	if !common.CheckNil(field.ID, field.Name) {
		return errors.New(common.Enums.ErrorMessages.LackOfInfo)
	}
	tempField, _ := GetFieldByID(field.ID)
	if field.IsExist() && field.Name != tempField.Name {
		//ensure the field name is not in the db
		return errors.New(common.Enums.ErrorMessages.RecordExist)
	}
	field.ModifiedDate = time.Now()
	err := common.Resources.Update(fieldCollection, common.MgoQry.Bson("_id", field.ID), field, false)
	return err
}

//IsExist check to see if the field name is already exist in the database.
func (field *Field) IsExist() bool {
	tempField, _ := GetFieldByName(field.Name)
	if common.CheckNil(tempField.Name) {
		return true
	}
	return false
}

//DeleteField deletes the selected field by Id
func DeleteField(id bson.ObjectId) error {
	return common.Resources.Delete(fieldCollection, common.MgoQry.Bson("_id", id), false)
}

//GetField returns a field according to the filter query
func GetField(query bson.M) (*Field, error) {
	var field Field
	err := common.Resources.Get(fieldCollection, &field, query, nil)
	return &field, err
}

//GetFields returns fields according to the filter query
func GetFields(query bson.M) (*[]Field, error) {
	var fields []Field
	err := common.Resources.GetAll(fieldCollection, &fields, query, nil)
	return &fields, err
}

//GetFieldByID returns field by id
func GetFieldByID(id bson.ObjectId) (*Field, error) {
	return GetField(common.MgoQry.Bson("_id", id))
}

//GetFieldByName returns field by name
func GetFieldByName(name string) (*Field, error) {
	return GetField(common.MgoQry.Bson("name", name))
}

//GetFieldsByType returns fields with the same type name
func GetFieldsByType(typeName string) (*[]Field, error) {
	return GetFields(common.MgoQry.Bson("type", typeName))
}
