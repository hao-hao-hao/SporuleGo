package common

import (
	"errors"
	"math/rand"
	"reflect"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

//CheckNil checks if the item are empty, it will return true if it is not nil
func CheckNil(items ...interface{}) (result bool) {
	result = true
	for _, item := range items {
		value := reflect.ValueOf(item)
		//return false if the value is not valid
		if !value.IsValid() {
			result = false
			return result
		}
		switch value.Kind() {
		case reflect.Slice, reflect.String, reflect.Map, reflect.Array:
			result = value.Len() > 0
		case reflect.Ptr, reflect.Interface:
			//recursive call itself to check the real value
			valueElem := value.Elem()
			result = CheckNil(valueElem)
		case reflect.Struct:
			//return true by default, struct is always not nil
		default:
			result = false
		}

		if !result {
			//return the result if the result is false
			return result
		}
	}
	return result
}

//EncryptPassword Hashes the password
func EncryptPassword(password string) (string, error) {
	newPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(newPassword), err
}

//VerifyPassword verify the password and stored encrypted password
func VerifyPassword(password, encryptedPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(encryptedPassword), []byte(password)) == nil
}

//GetError returns error message if error is not nil
func GetError(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}

//GenerateRandomString provides random strings
func GenerateRandomString() string {
	time.Sleep(1 * time.Nanosecond)
	rand.Seed(time.Now().UnixNano())
	return strconv.Itoa(rand.Intn(88888))
}

//StringToObjectID returns MongoDB Object ID
func StringToObjectID(id string) (objectID bson.ObjectId, err error) {
	defer func() {
		if errTemp := recover(); errTemp != nil {
			err = errors.New(Enums.ErrorMessages.PageNotFound)
		}
	}()
	objectID = bson.ObjectIdHex(id)
	return objectID, err
}

//QueryHelper is query helper for mongo DB
type QueryHelper struct{}

//MgoQry is a public exposed functions for buildding different querys
var MgoQry QueryHelper

//All will match all queies in arrary
func (query *QueryHelper) All(values ...interface{}) bson.M {
	return bson.M{"$all": values}
}

//In will match any queries in arrary
func (query *QueryHelper) In(values ...interface{}) bson.M {
	return bson.M{"$in": values}
}

//Nin will match anything other than the queies in arrary
func (query *QueryHelper) Nin(values ...interface{}) bson.M {
	return bson.M{"$nin": values}
}

//Eq matches equale comparison
func (query *QueryHelper) Eq(value interface{}) bson.M {
	return bson.M{"$eq": value}
}

//Gt matches greater comparison
func (query *QueryHelper) Gt(value interface{}) bson.M {
	return bson.M{"$gt": value}
}

//Gte matches greater or equal comparison
func (query *QueryHelper) Gte(value interface{}) bson.M {
	return bson.M{"$gte": value}
}

//Lt matches less comparison
func (query *QueryHelper) Lt(value interface{}) bson.M {
	return bson.M{"$lt": value}
}

//Lte matches less or equal comparison
func (query *QueryHelper) Lte(value interface{}) bson.M {
	return bson.M{"$lte": value}
}

//And provides and relationship
func (query *QueryHelper) And(values ...interface{}) bson.M {
	return bson.M{"$and": values}
}

//Or provides and relationship
func (query *QueryHelper) Or(values ...interface{}) bson.M {
	return bson.M{"$Or": values}
}

//Not provides NOT relationship
func (query *QueryHelper) Not(value interface{}) bson.M {
	return bson.M{"$not": value}
}

//Nor provides NOR relationship
func (query *QueryHelper) Nor(values ...interface{}) bson.M {
	return bson.M{"$nor": values}
}

//Slice sets the item skip and limit of the query
func (query *QueryHelper) Slice(skip, limit int) bson.M {
	return bson.M{"$slice": []int{skip, limit}}
}

//Select takes fields name and returns the "filenames":"1" to select the input fields
func (query *QueryHelper) Select(values ...string) bson.M {
	selector := bson.M{}
	for _, value := range values {
		selector[value] = 1
	}
	return selector
}

//Bson returns a bson.M key value
func (query *QueryHelper) Bson(key string, value interface{}) bson.M {
	return bson.M{key: value}
}
