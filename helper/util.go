package helper

import (
	"github.com/martinyonathann/bookingapp/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func DateTime(format string) string {
	currentTime := time.Now()
	result := currentTime.Format(format) //"2006-01-02 15:04:05"
	return result
}

func JwtGenerator(username, key string) string {
	//Generate Token JWT for auth
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["dateCreated"] = DateTime("2006-01-02")

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
