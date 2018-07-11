package models

import (
	"errors"
	"fmt"
	"sporule/api/app/modules/common"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//userCollection is the collection name for Model User in mongo db
const userCollection = "user"

//User is user account which will include authentications
type User struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	Email       string        `bson:"email,omitempty"`
	Password    string        `bson:"password,omitempty"`
	Name        string        `bson:"name,omitempty"`
	LastLogin   time.Time     `bson:"lastLogin,omitempty"`
	FailedLogin uint          `bson:"failedLogin,omitempty"`
	TokenSalt   string        `bson:"resetToken,omitempty"`
	IsDisabled  bool          `bson:"isDisabled,omitempty"`
	Roles       []Role        `bson:"roles,omitempty"`
}

//NewUser Constructor, It will create a new user object, and inject mongodb ID and hash password automatically
func NewUser(email, password, name string, roles []Role) (user *User, err error) {
	user = &User{}
	member, error := GetRoleByName(common.Enums.Roles.Admin)
	if error == nil {
		roles = append(roles, *member)
	}
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
	//flag change
	//cccc, _ := bson.Marshal(*user)
	//print(string(cccc[:]))
	//err := common.Resources.GetAll(userCollection, &users, bson.M{"roles.name": common.MgoQry.Nin("ABC", "BBC")})
	qry := common.MgoQry.And(common.MgoQry.Bson("roles.name", common.MgoQry.Nin("asdasd")))
	fmt.Printf("%v", qry)
	err := common.Resources.GetAll(userCollection, &users, qry, func(query *mgo.Query) *mgo.Query {
		return query.Select(common.MgoQry.Select("email"))
	})
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
