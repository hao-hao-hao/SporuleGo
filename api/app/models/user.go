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
	ResetToken  string        `bson:"resetToken"`
	IsDisabled  bool          `bson:"isDisabled"`
	Roles       []Role        `bson:"roles"`
}

//NewUser Constructor, It will inject mongodb ID automatically
func NewUser(email, password, name string, roles []Role) (user *User, err error) {
	//check if at least email, password, string and role is not nil
	isValid := common.CheckNil(email, password, name, roles)
	encryptedPassword, _ := common.EncryptPassword(password)
	if isValid {
		user = &User{
			Email:       email,
			Password:    encryptedPassword,
			Name:        name,
			FailedLogin: 0,
			IsDisabled:  false,
			Roles:       roles,
		}
	}
	return user, err
}

//Register adds User to database if it is not exist already. It will return an error if the user it is in the database
func (user *User) Register() (err error) {
	if common.CheckNil(user.Email, user.Password, user.Name) {
		tempUser, _ := GetUserByEmail(user.Email)
		if !common.CheckNil(tempUser.Email) {
			//add user if the email is not in database
			user.ID = bson.NewObjectId()
			err = common.Create(collection, user)
		} else {
			err = errors.New("Your Email Address is already exists")
		}
	} else {
		err = errors.New("Please ensure you have post at least Email, Password and Name")
	}
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
