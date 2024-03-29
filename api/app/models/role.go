package models

import (
	"errors"
	"sporule/api/app/modules/common"
	"time"

	"gopkg.in/mgo.v2/bson"
)

//roleCollection is the collection name for Model User in mongo db
const roleCollection = "role"

//Role is for permission management
type Role struct {
	ID           bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
	Name         string        `bson:"name" json:"name,omitempty"`
	CreatedDate  time.Time     `bson:"createdDate,omitempty" json:"createdDate,omitempty"`
	ModifiedDate time.Time     `bson:"modeifiedDate,omitempty" json:"modeifiedDate,omitempty"`
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
	if !common.CheckNil(role.ID, role.Name) {
		return errors.New(common.Enums.ErrorMessages.LackOfInfo)
	}
	tempRole, _ := GetRoleByID(role.ID)
	if role.IsExist() && role.Name != tempRole.Name {
		return errors.New(common.Enums.ErrorMessages.RecordExist)
	}
	role.CreatedDate = time.Now()
	role.ModifiedDate = time.Now()
	if common.Resources.Create(roleCollection, role) != nil {
		return errors.New(common.Enums.ErrorMessages.LackOfInfo)
	}
	return nil
}

//Update updates the role to the database
func (role *Role) Update() error {
	if !common.CheckNil(role.ID, role.Name) {
		return errors.New(common.Enums.ErrorMessages.LackOfInfo)
	}
	tempRole, _ := GetRoleByID(role.ID)
	if role.IsExist() && role.Name != tempRole.Name {
		//need to ensure the new role name is not exist in the db
		return errors.New(common.Enums.ErrorMessages.RecordExist)
	}
	role.ModifiedDate = time.Now()
	err := common.Resources.Update(roleCollection, common.MgoQry.Bson("_id", role.ID), role, false)
	return err
}

//IsExist check to see if the role name is already exist in the database.
func (role *Role) IsExist() bool {
	tempRole, _ := GetRoleByName(role.Name)
	if common.CheckNil(tempRole.Name) {
		return true
	}
	return false
}

//DeleteRoleByID deletes the selected role by Id
func DeleteRoleByID(id bson.ObjectId) error {
	return common.Resources.Delete(roleCollection, common.MgoQry.Bson("_id", id), true)
}

//GetRole returns a role according to the filter query
func GetRole(query bson.M) (*Role, error) {
	var role Role
	err := common.Resources.Get(roleCollection, &role, query, nil)
	return &role, err
}

//GetRoles returns roles according to the filter query
func GetRoles(query bson.M) (*[]Role, error) {
	var roles []Role
	err := common.Resources.GetAll(roleCollection, &roles, query, nil)
	return &roles, err
}

//GetRoleByID returns field by id
func GetRoleByID(id bson.ObjectId) (*Role, error) {
	role, err := GetRole(common.MgoQry.Bson("_id", id))
	return role, err
}

//GetRoleByName returns field by id
func GetRoleByName(name string) (*Role, error) {
	role, err := GetRole(common.MgoQry.Bson("name", name))
	return role, err
}
