package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/martinyonathann/bookingapp/db"
	"github.com/martinyonathann/bookingapp/helper"
	"github.com/martinyonathann/bookingapp/models"
)

const (
	QUERY_INSERT_HOTEL = "INSERT INTO hotels (hotel_name, hotel_address, hotel_city, hotel_state, hotel_zip, hotel_country, hotel_price) VALUES($1, $2, $3, $4, $5, $6, $7)RETURNING hotel_id;"
	QUERY_DELETE_HOTEL = "DELETE FROM hotels WHERE hotel_id = $1;"
	QUERY_SELECT_HOTEL = "SELECT * FROM hotels"
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

func DeleteHotel(c echo.Context) error {
	c.Response().Header().Set("Content-Type", "application/json")
	id, _ := strconv.Atoi(c.Param("id"))
	var count int64
	DB := db.DB()
	res, _ := DB.Exec(QUERY_DELETE_HOTEL, id)
	count, _ = res.RowsAffected()
	if count < 1 {
		resp := c.JSON(http.StatusNotFound, models.ResponseFailed(http.StatusNotFound, "HOTEL NOT FOUND"))
		return resp
	}
	resp := c.JSON(http.StatusOK, models.ResponseSuccess(nil, "SUCCESS DELETE HOTEL"))
	return resp
}
func GetAllHotel(c echo.Context) error {
	c.Response().Header().Set("Content-Type", "application/json")

	DB := db.DB()
	var hotel models.Hotel

	rows, err := DB.Query(QUERY_SELECT_HOTEL)
	if err != nil {
		log.Error(err.Error())
		resp := c.JSON(http.StatusNotFound, models.ResponseFailed(http.StatusNotFound, err.Error()))
		return resp
	}
	defer rows.Close()
	var hotels []models.Hotel
	for rows.Next() {
		rows.Scan(&hotel.HotelId, &hotel.Name, &hotel.Address, &hotel.City, &hotel.State, &hotel.Zip, &hotel.Country, &hotel.Price)
		hotels = append(hotels, hotel)
	}
	resp := c.JSON(http.StatusOK, models.ResponseSuccess(hotels, "SUCCESS"))
	return resp
}
