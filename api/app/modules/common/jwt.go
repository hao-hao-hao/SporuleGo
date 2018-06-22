package common

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"

	jwt "github.com/dgrijalva/jwt-go"
)

//AuthenticationClaims is the claims struct for jwt token
type AuthenticationClaims struct {
	Email string `json:"email"`
	Salt  string `json:"salt"`
	jwt.StandardClaims
}

//GenerateJWT generates JSON WEB TOKEN for authentication by using email address
func GenerateJWT(email string, salt string) (string, error) {
	claims := AuthenticationClaims{
		email, salt,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * Config.JWTLife).Unix(),
			Issuer:    Config.JWTIssuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(Config.JWTKey))
}

//VerifyToken verify and parse the token to email address
func VerifyToken(tokenString string) (string, string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &AuthenticationClaims{}, keyfunc)
	if err == nil {
		claims := token.Claims.(*AuthenticationClaims)
		return claims.Email, claims.Salt, err
	}
	return "", "", err
}

//keyfunc is the function that validates the sign method and then return the JWT Key
func keyfunc(token *jwt.Token) (interface{}, error) {
	if jwt.SigningMethodHS256 != token.Method {
		return nil, errors.New("Invalid signing algorithm")
	}
	return []byte(Config.JWTKey), nil
}

//SetIDInHeader sets the user id(can be email) in the header
func SetIDInHeader(c *gin.Context, id string) {
	c.Set(Enums.Others.IDInHeader, id)
}

//GetIDInHeader gets the user id(can be email) in the header, it will return 401 if the id is empty
func GetIDInHeader(c *gin.Context) string {
	id, _ := c.Get(Enums.Others.IDInHeader)
	if !CheckNil(id) {
		//return 401 if id is not found to be defensive
		HTTPResponse401(c)
		return ""
	}
	return id.(string)
}
