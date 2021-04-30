package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gomodule/redigo/redis"
	"github.com/martinyonathann/bookingapp/db"
	"github.com/martinyonathann/bookingapp/helper"
	"github.com/martinyonathann/bookingapp/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
)

const (
	QUERY_COUNT_REGISTRATION  = "SELECT COUNT(*) FROM users WHERE username = $1;"
	QUERY_INSERT_REGISTRATION = "INSERT INTO users(username, first_name, last_name, password, date_created) VALUES ($1, $2, $3, $4, $5) RETURNING user_id;"
	QUERY_SELECT_LOGIN        = "SELECT username, first_name, last_name, password, date_created FROM users WHERE username = $1"
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
	user.DateCreated = helper.DateTime("2006-01-02 15:04:05")

	// stmt, err := DB.Prepare(QUERY_INSERT_REGISTRATION)
	// if err != nil {
	// 	resp := c.JSON(http.StatusInternalServerError, helper.ErrorLog(http.StatusInternalServerError, "Error when prepare statement : "+err.Error(), "EXT_REF"))
	// 	return resp
	// }
	err = DB.QueryRow(QUERY_INSERT_REGISTRATION, user.Username, user.Firstname, user.Lastname, user.Password, user.DateCreated).Scan(&user.UserId)

	if err != nil {
		log.Error(err)
		resp := c.JSON(http.StatusInternalServerError, helper.ErrorLog(http.StatusInternalServerError, "Error when execute statement : "+err.Error(), "EXT_REF"))
		return resp
	}

	user.Password = ""
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

	//Validate username
	err = DB.QueryRow(sqlStatement, user.Username).Scan(&result.Username, &result.Firstname, &result.Lastname, &result.Password, &result.DateCreated)
	if err != nil {
		log.Error(err)
		resp := c.JSON(http.StatusInternalServerError, helper.ErrorLog(http.StatusInternalServerError, "Invalid Username", "EXT_REF"))
		return resp
	}

	//Hashing and compare password
	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(user.Password))
	if err != nil {
		log.Error("Invalid Password :", err)
		resp := c.JSON(http.StatusInternalServerError, helper.ErrorLog(http.StatusInternalServerError, "Invalid Password", "EXT_REF"))
		return resp
	}

	result.Token = helper.JwtGenerator(result.Username, "secret")
	result.Password = ""

	//Redis
	var cacheName string = result.Username + helper.DateTime("2006-01-02")
	conn, err := redis.Dial("tcp", "localhost:6379")

	if err != nil {
		log.Error("Error when connect redis :", err)
		resp := c.JSON(http.StatusInternalServerError, helper.ErrorLog(http.StatusInternalServerError, "Error when connect redis", "EXT_REF"))
		return resp
	}
	jsonData, _ := json.Marshal(&result)
	_, err = conn.Do("SET", cacheName, string(jsonData))
	if err != nil {
		log.Panic(err)
	}

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
		//Connection to redis for get data user.
		var cacheName string = claims["username"].(string) + helper.DateTime("2006-01-02")
		conn, err := redis.Dial("tcp", "localhost:6379")
		if err != nil {
			log.Error("Error when connect redis :", err)
			resp := c.JSON(http.StatusInternalServerError, helper.ErrorLog(http.StatusInternalServerError, "Error when connect redis", "EXT_REF"))
			return resp
		}

		//Get data from cache
		reply, err := redis.Bytes(conn.Do("GET", cacheName))
		if err != nil {
			log.Error("Error when get data redis :", err)
			resp := c.JSON(http.StatusInternalServerError, helper.ErrorLog(http.StatusInternalServerError, "Error when get data redis", "EXT_REF"))
			return resp
		}
		//Umarshal data redis to model result.
		err = json.Unmarshal(reply, &result)
		if err != nil {
			log.Error("error when Unmarshal data cache", err.Error())
			resp := c.JSON(http.StatusInternalServerError, helper.ErrorLog(http.StatusInternalServerError, "error when Unmarshal data cache", "EXT_REF"))
			return resp
		}

		log.Info()
		resp := c.JSON(http.StatusOK, result)
		return resp
	} else {
		log.Error(err.Error())
		resp := c.JSON(http.StatusInternalServerError, helper.ErrorLog(http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", "EXT_REF"))
		return resp
	}

}
