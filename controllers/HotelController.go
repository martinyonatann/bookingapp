package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/martinyonathann/bookingapp/db"
	"github.com/martinyonathann/bookingapp/helper"
	"github.com/martinyonathann/bookingapp/models"
)

const (
	QUERY_INSERT_HOTEL = "INSERT INTO hotels (hotel_name, hotel_address, hotel_city, hotel_state, hotel_zip, hotel_country, hotel_price) VALUES($1, $2, $3, $4, $5, $6, $7)RETURNING hotel_id;"
)

func AddHotel(c echo.Context) error {
	c.Response().Header().Set("Content-Type", "application/json")

	var hotel models.Hotel
	body, _ := ioutil.ReadAll(c.Request().Body)
	err := json.Unmarshal(body, &hotel)

	if err != nil {
		log.Error(err.Error())
		resp := c.JSON(http.StatusInternalServerError, helper.ErrorLog(http.StatusInternalServerError, "Error when Unmarshal body", "EXT_REF"))
		return resp
	}
	DB := db.DB()
	err = DB.QueryRow(QUERY_INSERT_HOTEL, hotel.Name, hotel.Address, hotel.City, hotel.State, hotel.Zip, hotel.Country, hotel.Price).Scan(&hotel.HotelId)
	if err != nil {
		log.Error(err.Error())
		resp := c.JSON(http.StatusInternalServerError, helper.ErrorLog(http.StatusInternalServerError, "Failed while try to execute query", "EXT_REF"))
		return resp
	}
	return c.JSON(http.StatusAccepted, hotel)
}
