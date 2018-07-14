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

//QueryHelper is query helper struct for mongo DB, this is purely for better organisations
type QueryHelper struct{}

//MgoQry is a public exposed functions for buildding different querys
var MgoQry QueryHelper

//Bson returns a bson.M key value pair
func (query *QueryHelper) Bson(key string, value interface{}) bson.M {
	return bson.M{key: value}
}

//Bsons returns multiple bson.M key value pairs
func (query *QueryHelper) Bsons(keyValuePairs map[string]interface{}) bson.M {
	qry := bson.M{}
	for key, value := range keyValuePairs {
		qry[key] = value
	}
	return qry
}

//All will match all queies in arrary
func (query *QueryHelper) All(values ...interface{}) bson.M {
	return query.Bson("$all", values)
}

//In will match any queries in arrary
func (query *QueryHelper) In(values ...interface{}) bson.M {
	return query.Bson("$in", values)
}

//Nin will match anything other than the queies in arrary
func (query *QueryHelper) Nin(values ...interface{}) bson.M {
	return query.Bson("$nin", values)
}

//Eq matches equale comparison
func (query *QueryHelper) Eq(value interface{}) bson.M {
	return query.Bson("$eq", value)
}

//Gt matches greater comparison
func (query *QueryHelper) Gt(value interface{}) bson.M {
	return query.Bson("$gt", value)
}

//Gte matches greater or equal comparison
func (query *QueryHelper) Gte(value interface{}) bson.M {
	return query.Bson("$gte", value)
}

//Lt matches less comparison
func (query *QueryHelper) Lt(value interface{}) bson.M {
	return query.Bson("$lt", value)
}

//Lte matches less or equal comparison
func (query *QueryHelper) Lte(value interface{}) bson.M {
	return query.Bson("$lte", value)
}

//And provides and relationship
func (query *QueryHelper) And(queries ...interface{}) bson.M {
	return query.Bson("$and", queries)
}

//Or provides and relationship
func (query *QueryHelper) Or(values ...interface{}) bson.M {
	return query.Bson("$Or", values)
}

//Not provides NOT relationship
func (query *QueryHelper) Not(value interface{}) bson.M {
	return query.Bson("$not", value)
}

//Nor provides NOR relationship
func (query *QueryHelper) Nor(values ...interface{}) bson.M {
	return query.Bson("$nor", values)
}

//Slice sets the item skip and limit of the query
func (query *QueryHelper) Slice(skip, limit int) bson.M {
	return query.Bson("$slice", []int{skip, limit})
}

//Select takes fields name and returns the "filenames":"1" to select the input fields
func (query *QueryHelper) Select(isSelect bool, values ...string) bson.M {
	selector := bson.M{}
	for _, value := range values {
		selector[value] = isSelect
	}
	return selector
}

//Match returns the bson.M for $match operation, this is for aggregation queries only
func (query *QueryHelper) Match(qry interface{}) bson.M {
	return query.Bson("$match", qry)
}

//Limit returns the bson.M for $limit operation, this is for aggregation queries only
func (query *QueryHelper) Limit(maxReturn int) bson.M {
	return query.Bson("$limit", maxReturn)
}

//Sort returns the bson.M for $sort operation, it takes only one fields, this is for aggregation queries only
func (query *QueryHelper) Sort(field string, isDescending bool) bson.M {
	order := 1
	if isDescending {
		// 1 is ascending and -1 is descending
		order = -1
	}
	return query.Bson("$sort", query.Bson(field, order))
}

//Sorts returns the bson.M for $sort operation, it takes multiple key value pairs, this is for aggregation queries only
//Use field name as key and 1/-1 in the value,1 is ascending and -1 is descending
func (query *QueryHelper) Sorts(keyValuePairs map[string]interface{}) bson.M {
	return query.Bson("$sort", query.Bsons(keyValuePairs))
}

//Project returns the bson.M for $project operation which sets the selected fields in SQL, this is for aggregation queries only
func (query *QueryHelper) Project(qry interface{}) bson.M {
	return query.Bson("$project", qry)
}

//LookUp returns the bson.M for $lookup operation, this is for aggregation queries only
func (query *QueryHelper) LookUp(from, localField, foreignField, as string) bson.M {
	qry := make(map[string]interface{})
	qry["from"] = from
	qry["localField"] = localField
	qry["foreignField"] = foreignField
	qry["as"] = as
	return query.Bson("$lookup", query.Bsons(qry))
}
