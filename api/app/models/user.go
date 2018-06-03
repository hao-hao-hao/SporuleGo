package models

import (
	"errors"
	"sporule/api/app/modules/common"
	"time"

	"gopkg.in/mgo.v2/bson"
)

//UserCollection is the collection name for Model User in mongo db
const collection = "user"

//User is user account which will include authentications
type User struct {
	ID          bson.ObjectId `bson:"_id"`
	Email       string        `bson:"email"`
	Password    string        `bson:"password"`
	Name        string        `bson:"name"`
	LastLogin   time.Time     `bson:"lastLogin"`
	FailedLogin uint          `bson:"failedLogin"`
	TokenSalt   string        `bson:"resetToken"`
	IsDisabled  bool          `bson:"isDisabled"`
	Roles       []Role        `bson:"roles"`
}

//NewUser Constructor, It will inject mongodb ID and hash password automatically
func (user *User) NewUser(email, password, name string, roles []Role) (err error) {
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
		user.Roles = roles
		user.TokenSalt = common.GenerateRandomString()
	} else {
		err = errors.New(common.Enums.ErrorMessages.LackOfRegInfo)
	}
	return err
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
	if user.NewUser(user.Email, user.Password, user.Name, nil) != nil {
		return errors.New(common.Enums.ErrorMessages.SystemError)
	}
	if common.Create(collection, user) != nil {
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

//UpdateTokenSalt updates the token salt to invalid the user id
func (user *User) UpdateTokenSalt() error {
	user.TokenSalt = common.GenerateRandomString()
	return user.UpdateUser()
}

//UpdateUser updates the user to the database
func (user *User) UpdateUser() error {
	err := common.Update(collection, bson.M{"email": user.Email}, user, false)
	return err
}

//GetUser returns a user according to the filter query
func GetUser(query bson.M) (*User, error) {
	var user User
	s, c := common.Collection(collection)
	defer s.Close()
	err := c.Find(query).One(&user)
	return &user, err
}

//GetUsers returns an user slice according to the filter
func GetUsers(query bson.M) (*[]User, error) {
	var users []User
	s, c := common.Collection(collection)
	defer s.Close()
	err := c.Find(query).All(&users)
	return &users, err
}

//GetUserByEmail returns user by giving the email address
func GetUserByEmail(email string) (*User, error) {
	user, err := GetUser(bson.M{"email": email})
	return user, err
}

//GetUserByID returns user by giving the user ID
func GetUserByID(ID string) (*User, error) {
	user, err := GetUser(bson.M{"ID": ID})
	return user, err
}
