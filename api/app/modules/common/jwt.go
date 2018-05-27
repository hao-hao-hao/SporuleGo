package common

import (
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

//AuthenticationClaims is the claims struct for jwt token
type AuthenticationClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

//GenerateJWT generates JSON WEB TOKEN for authentication by using email address
func GenerateJWT(email string) (string, error) {
	claims := AuthenticationClaims{
		email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * Config.JWTLife).Unix(),
			Issuer:    Config.JWTIssuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(Config.JWTKey))
}

//VerifyToken verify and parse the token to email address
func VerifyToken(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &AuthenticationClaims{}, keyfunc)
	if err == nil {
		claims := token.Claims.(*AuthenticationClaims)
		return claims.Email, err
	}
	return "", err
}

//keyfunc is the function that validates the sign method and then return the JWT Key
func keyfunc(token *jwt.Token) (interface{}, error) {
	if jwt.SigningMethodHS256 != token.Method {
		return nil, errors.New("Invalid signing algorithm")
	}
	return []byte(Config.JWTKey), nil
}
