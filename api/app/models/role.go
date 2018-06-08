package models

import (
	"errors"
	"sporule/api/app/modules/common"

	"gopkg.in/mgo.v2/bson"
)

//roleCollection is the collection name for Model User in mongo db
const roleCollection = "role"

//Role is for permission management
type Role struct {
	ID   bson.ObjectId `bson:"ID"`
	Name string        `bson:"name"`
}

//NewRole is the constructor for Role
func NewRole(name string) (*Role, error) {
	if !common.CheckNil(name) {
		return nil, errors.New(common.Enums.ErrorMessages.LackOfInfo)
	}
	role := &Role{}
	role.ID = bson.NewObjectId()
	role.Name = name
	return role, nil
}

//Insert inserts the role to the database
func (role *Role) Insert() error {
	if !common.CheckNil(role.Name) {
		return errors.New(common.Enums.ErrorMessages.LackOfInfo)
	}
	tempRole, _ := GetUserByEmail(role.Name)
	if common.CheckNil(tempRole.Name) {
		return errors.New(common.Enums.ErrorMessages.RecordExist)
	}
	if common.Create(roleCollection, role) != nil {
		return errors.New(common.Enums.ErrorMessages.LackOfInfo)
	}
	return nil
}

//Update updates the role to the database
func (role *Role) Update() error {
	err := common.Update(roleCollection, bson.M{"_id": role.ID}, role, false)
	return err
}

//DeleteRole deletes the selected role by Id
func DeleteRole(id bson.ObjectId) error {
	return common.Delete(roleCollection, bson.M{"_id": id}, true)
}

//GetRole returns a role according to the filter query
func GetRole(query bson.M) (*Role, error) {
	var role Role
	s, c := common.Collection(roleCollection)
	defer s.Close()
	err := c.Find(query).One(&role)
	return &role, err
}

//GetRoleByID returns field by id
func GetRoleByID(id bson.ObjectId) (*Role, error) {
	role, err := GetRole(bson.M{"_id": id})
	return role, err
}

//GetRoleByName returns field by id
func GetRoleByName(name string) (*Role, error) {
	role, err := GetRole(bson.M{"name": name})
	return role, err
}
