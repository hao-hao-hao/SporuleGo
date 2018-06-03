package common

import (
	"math/rand"
	"reflect"
	"strconv"

	"golang.org/x/crypto/bcrypt"
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
	} else {
		return err.Error()
	}
}

//GenerateRandomString provides random strings
func GenerateRandomString() string {
	return strconv.Itoa(rand.Intn(88888))
}
