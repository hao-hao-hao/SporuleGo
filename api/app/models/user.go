package models

import (
	"errors"
	"sporule/api/app/modules/common"
	"time"

	"gopkg.in/mgo.v2/bson"
)

//userCollection is the collection name for Model User in mongo db
const userCollection = "user"

//User is user account struct
type User struct {
	ID               bson.ObjectId          `bson:"_id,omitempty" json:"_id,omitempty"`
	Email            string                 `bson:"email,omitempty" json:"email,omitempty"`
	Password         string                 `bson:"password,omitempty" json:"-"`
	RoleIds          []bson.ObjectId        `bson:"roleIds,omitempty" json:"-"`
	Roles            []Role                 `bson:"roles,omitempty" json:"roles,omitempty"`
	TokenSalt        string                 `bson:"tokenSalt,omitempty" json:"-"`
	LastAccess       time.Time              `bson:"lastAccess,omitempty" json:"lastAccess,omitempty"`
	FailedLogin      int                    `bson:"failedLogin,omitempty" json:"failedLogin,omitempty"`
	IsDisabled       bool                   `bson:"isDisabled,omitempty" json:"isDisabled,omitempty"`
	CustomAttributes map[string]interface{} `bson:"customAttributes,omitempty" json:"customAttributes,omitempty"`
}

//NewUser Constructor, It will create a new user object, and inject mongodb ID and hash password automatically
func NewUser(email, password string) (user *User, err error) {
	user = &User{}
	//check if at least email, password is not nil
	isValid := common.CheckNil(email, password)
	if isValid {
		encryptedPassword, _ := common.EncryptPassword(password)
		user.ID = bson.NewObjectId()
		user.Email = email
		user.Password = encryptedPassword
		user.FailedLogin = 0
		user.IsDisabled = false
		user.TokenSalt = common.GenerateRandomString()
	} else {
		err = errors.New(common.Enums.ErrorMessages.LackOfRegInfo)
		return nil, err
	}
	return user, nil
}

//Register adds User to database if it is not exist already. It will return an error if the user it is in the database.
//We only expect user without roles to use register function
func (user *User) Register() error {
	if !common.CheckNil(user.ID, user.Email, user.Password) {
		return errors.New(common.Enums.ErrorMessages.LackOfRegInfo)
	}
	tempUser, _ := GetUserByEmail(user.Email)
	if common.CheckNil(tempUser.Email) {
		return errors.New(common.Enums.ErrorMessages.UserExist)
	}
	//get the default role
	basicRole, err := GetRoleByName(common.Config.BasicMember)
	if err != nil {
		return err
	}
	user.RoleIds = append(user.RoleIds, basicRole.ID)
	//remove all User Roles as we only use RoleIDs in the database
	user.Roles = nil
	if common.Resources.Create(userCollection, user) != nil {
		return errors.New(common.Enums.ErrorMessages.SystemError)
	}
	return nil
}

//VerifyUser verifies the email and password, it also return the user object
func VerifyUser(email, password string) (*User, error) {
	if !common.CheckNil(email, password) {
		return nil, errors.New(common.Enums.ErrorMessages.AuthFailed)
	}
	user, err := GetUserByEmail(email)
	if err != nil {
		return nil, errors.New(common.Enums.ErrorMessages.AuthFailed)
	}
	if !common.VerifyPassword(password, user.Password) || user.IsDisabled {
		//add failed login
		user.FailedLogin++
		if user.FailedLogin >= common.Config.MaximumFailedLogin {
			//disable user when it exceed the maximum failed login
			user.IsDisabled = true
		}
		//update user to the db
		user.Update(user.ID)
		return nil, errors.New(common.Enums.ErrorMessages.AuthFailed)
	}
	return user, nil
}

//UpdateTokenSalt updates the token salt of the user to invalid the user token
func (user *User) UpdateTokenSalt() error {
	user.TokenSalt = common.GenerateRandomString()
	return user.Update(user.ID)
}

//Update updates the user to the database
func (user *User) Update(id bson.ObjectId) error {
	//remove all User Roles as we only use RoleIDs in the database
	user.Roles = nil
	err := common.Resources.Update(userCollection, bson.M{"_id": id}, user, false)
	return err
}

//ChangeDisableStatus will change user IsDisabled property from true to false or false to true
func (user *User) ChangeDisableStatus() error {
	if user.IsDisabled {
		user.IsDisabled = false
	} else {
		user.IsDisabled = true
	}
	return user.Update(user.ID)
}

//GetUser returns a user according to the filter query
func GetUser(filter bson.M) (*User, error) {
	var user User
	err := common.Resources.AggGetAll(userCollection, &user,
		common.MgoQry.LookUp("role", "roleIds", "_id", "roles"),
		common.MgoQry.Match(filter),
		common.MgoQry.Limit(1))
	return &user, err
}

//GetUsers returns an user slice according to the filter
func GetUsers(filter bson.M) (*[]User, error) {
	var users []User
	err := common.Resources.AggGetAll(userCollection, &users,
		common.MgoQry.LookUp("role", "roleIds", "_id", "roles"),
		common.MgoQry.Match(filter))
	return &users, err
}

//GetUserByEmail returns user by giving the email address
func GetUserByEmail(email string) (*User, error) {
	return GetUser(common.MgoQry.Bson("email", email))
}

//GetUserByID returns user by giving the user ID
func GetUserByID(ID bson.ObjectId) (*User, error) {
	return GetUser(common.MgoQry.Bson("_id", ID))
}

//GetDisabledUsers returns a list of disabled users
func GetDisabledUsers() (*[]User, error) {
	return GetUsers(common.MgoQry.Bson("isDisabled", true))
}

//GetUsersByRole returns a list of user with the role name
func GetUsersByRole(roleName string) (*[]User, error) {
	return GetUsers(common.MgoQry.Bson("roles.name", roleName))
}

//GetUsersGeneral returns user by using aggregate queries
func GetUsersGeneral(ID bson.ObjectId, email string, roleName string) (*[]User, error) {
	filters := []bson.M{}
	//Apply filters if the filter value is not nil
	if common.CheckNil(ID) {
		filters = append(filters, common.MgoQry.Bson("_id", ID))
	}
	if common.CheckNil(email) {
		filters = append(filters, common.MgoQry.Bson("email", email))
	}
	if common.CheckNil(roleName) {
		filters = append(filters, common.MgoQry.Bson("roles.name", roleName))
	}
	return GetUsers(common.MgoQry.And(filters))
}
