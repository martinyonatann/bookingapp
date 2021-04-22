package controllers

import (
	"bookingapp/db"
	"bookingapp/helper"
	"bookingapp/models"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
)

const (
	QUERY_COUNT_REGISTRATION  = "SELECT COUNT(*) FROM users WHERE username = $1;"
	QUERY_INSERT_REGISTRATION = "INSERT INTO users(username, first_name, last_name, password, token, date_created) VALUES ($1, $2, $3, $4, $5, $6);"
	QUERY_SELECT_LOGIN        = "SELECT username, first_name, last_name, password, token, date_created FROM users WHERE username = $1"
)

func Registration(c echo.Context) error {
	c.Response().Header().Set("Content-Type", "application/json")
	var user models.Users
	body, _ := ioutil.ReadAll(c.Request().Body)
	err := json.Unmarshal(body, &user)

	if err != nil {
		return err
	}
	DB := db.DB()

	var count int

	sqlStatement := QUERY_COUNT_REGISTRATION
	err = DB.QueryRow(sqlStatement, user.Username).Scan(&count)
	if err != nil {
		return err
	}
	if count > 0 {
		log.Error("username already used")
		resp := c.JSON(http.StatusConflict, helper.ErrorLog(http.StatusConflict, "username already used", "EXT_REF"))
		return resp
	}

	//hashing password
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 5)
	if err != nil {
		resp := c.JSON(http.StatusInternalServerError, helper.ErrorLog(http.StatusInternalServerError, " Error While Hashing Password", "EXT_REF"))
		return resp
	}

	user.Password = string(hash)
	user.DateCreated = helper.DateTime()
	user.Token = helper.JwtGenerator(user.Username, "secretkey")

	stmt, err := DB.Prepare(QUERY_INSERT_REGISTRATION)
	if err != nil {
		resp := c.JSON(http.StatusInternalServerError, helper.ErrorLog(http.StatusInternalServerError, "Error when prepare statement : "+err.Error(), "EXT_REF"))
		return resp
	}

	_, err = stmt.Exec(user.Username, user.Firstname, user.Lastname, user.Password, user.Token, user.DateCreated)
	if err != nil {
		log.Error(err)
		resp := c.JSON(http.StatusInternalServerError, helper.ErrorLog(http.StatusInternalServerError, "Error when execute statement : "+err.Error(), "EXT_REF"))
		return resp
	}

	// //logging needed
	// b, _ := json.Marshal(user)
	// fmt.Print("response : ", string(b))

	resp := c.JSON(http.StatusOK, user)
	log.Info()
	return resp
}

func LoginController(c echo.Context) error {
	c.Response().Header().Set("Content-Type", "application/json")
	var user models.Users
	payload, _ := ioutil.ReadAll(c.Request().Body)
	err := json.Unmarshal(payload, &user)

	if err != nil {
		log.Error(err)
		return err
	}
	var result models.Users
	DB := db.DB()
	sqlStatement := QUERY_SELECT_LOGIN

	err = DB.QueryRow(sqlStatement, user.Username).Scan(&result.Username, &result.Firstname, &result.Lastname, &result.Password, &result.Token, &result.DateCreated)
	if err != nil {
		log.Error(err)
		resp := c.JSON(http.StatusInternalServerError, helper.ErrorLog(http.StatusInternalServerError, "Invalid Username", "EXT_REF"))
		return resp
	}

	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(user.Password))
	if err != nil {
		log.Error("Invalid Password :", err)
		resp := c.JSON(http.StatusInternalServerError, helper.ErrorLog(http.StatusInternalServerError, "Invalid Password", "EXT_REF"))
		return resp
	}

	result.Token = helper.JwtGenerator(result.Username, "secret")

	// //logging needed
	// b, _ := json.Marshal(result)
	// fmt.Print("response : ", string(b))

	//resp
	resp := c.JSON(http.StatusOK, result)
	log.Info()
	return resp
}

func ProfileHandler(c echo.Context) error {
	c.Response().Header().Set("Content-Type", "application/json")
	tokenString := c.Request().Header.Get("Authorization")
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	var result models.Users
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		result.Username = claims["username"].(string)
		result.Firstname = claims["firstname"].(string)
		result.Lastname = claims["lastname"].(string)

		log.Info()
		resp := c.JSON(http.StatusOK, result)
		return resp
	} else {
		log.Error(err.Error())
		resp := c.JSON(http.StatusInternalServerError, helper.ErrorLog(http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", "EXT_REF"))
		return resp
	}

}
