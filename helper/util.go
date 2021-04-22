package helper

import (
	"bookingapp/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func DateTime() string {
	currentTime := time.Now()
	result := currentTime.Format("2006-01-02 15:04:05") //yyyy-mm-dd HH:mm:ss
	return result
}

func JwtGenerator(username, key string) string {
	//Generate Token JWT for auth
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		return err.Error()
	}
	return tokenString
}

func ErrorLog(rc int, detail, ext_ref string) models.Error {
	var error models.Error
	error.ResponseCode = rc
	error.Message = "Failed"
	error.Detail = detail
	error.ExternalReference = ext_ref

	return error
}
