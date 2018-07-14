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
	ID          bson.ObjectId   `bson:"_id,omitempty" json:"_id"`
	Email       string          `bson:"email,omitempty" json:"email"`
	Password    string          `bson:"password,omitempty" json:"-"`
	Name        string          `bson:"name,omitempty" json:"name"`
	RoleIds     []bson.ObjectId `bson:"roleIds,omitempty" json:"-"`
	Roles       []Role          `bson:"roles,omitempty" json:"roles"`
	TokenSalt   string          `bson:"resetToken,omitempty" json:"-"`
	LastAccess  time.Time       `bson:"lastAccess,omitempty" json:"lastAccess"`
	FailedLogin uint            `bson:"failedLogin,omitempty" json:"-"`
	IsDisabled  bool            `bson:"isDisabled,omitempty" json:"isDisabled"`
}

//NewUser Constructor, It will create a new user object, and inject mongodb ID and hash password automatically
func NewUser(email, password, name string, roleIds []bson.ObjectId) (user *User, err error) {
	user = &User{}

	//check if at least email, password, string and role is not nil
	isValid := common.CheckNil(email, password, name)
	encryptedPassword, _ := common.EncryptPassword(password)
	if isValid {
		user.ID = bson.NewObjectId()
		user.Email = email
		user.Password = encryptedPassword
		user.Name = name
		user.FailedLogin = 0
		user.IsDisabled = false
		user.RoleIds = roleIds
		user.TokenSalt = common.GenerateRandomString()
	} else {
		err = errors.New(common.Enums.ErrorMessages.LackOfRegInfo)
		return nil, err
	}
	return user, nil
}

//Register adds User to database if it is not exist already. It will return an error if the user it is in the database
func (user *User) Register() error {
	if !common.CheckNil(user.Email, user.Password, user.Name) {
		return errors.New(common.Enums.ErrorMessages.LackOfRegInfo)
	}
	tempUser, _ := GetUserByEmail(user.Email)
	if common.CheckNil(tempUser.Email) {
		return errors.New(common.Enums.ErrorMessages.UserExist)
	}
	if common.Resources.Create(userCollection, user) != nil {
		return errors.New(common.Enums.ErrorMessages.SystemError)
	}
	return nil
}

//Verify verifies the user to see if it is valid
func (user *User) Verify() (err error) {
	if !common.CheckNil(user.Email, user.Password) {
		return errors.New(common.Enums.ErrorMessages.AuthFailed)
	}
	dbUser, err := GetUserByEmail(user.Email)
	if err != nil {
		return errors.New(common.Enums.ErrorMessages.AuthFailed)
	}
	if !common.VerifyPassword(user.Password, dbUser.Password) {
		return errors.New(common.Enums.ErrorMessages.AuthFailed)
	}
	user.TokenSalt = dbUser.TokenSalt
	return nil
}

//UpdateTokenSalt updates the token salt of the user to invalid the user token
func (user *User) UpdateTokenSalt() error {
	user.TokenSalt = common.GenerateRandomString()
	return user.Update(user.ID)
}

//Update updates the user to the database
func (user *User) Update(id bson.ObjectId) error {
	user.Roles = nil
	err := common.Resources.Update(userCollection, bson.M{"_id": id}, user, false)
	return err
}

//GetUser returns a user according to the filter query
func GetUser(query bson.M) (*User, error) {
	var user User
	err := common.Resources.Get(userCollection, &user, query, nil)
	return &user, err
}

//GetUsers returns an user slice according to the filter
func GetUsers(query bson.M) (*[]User, error) {
	var users []User
	err := common.Resources.GetAll(userCollection, &users, query, nil)
	return &users, err
}

//GetUsersA Updates the users
func GetUsersA() (*[]User, error) {
	var users []User
	filter := common.MgoQry.And(common.MgoQry.Bson("roles.name", "Admin"))
	err := common.Resources.AggGetAll(userCollection, &users,
		common.MgoQry.LookUp("role", "roleIds", "_id", "roles"),
		common.MgoQry.Match(filter),
		common.MgoQry.Project(common.MgoQry.Select(true, "email", "_id")))
	return &users, err
}

//GetUserByEmail returns user by giving the email address
func GetUserByEmail(email string) (*User, error) {
	user, err := GetUser(bson.M{"email": email})
	return user, err
}

//GetUserByID returns user by giving the user ID
func GetUserByID(ID bson.ObjectId) (*User, error) {
	user, err := GetUser(bson.M{"_id": ID})
	return user, err
}
