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
	"github.com/martinyonathann/bookingapp/helper/logging"

	"github.com/martinyonathann/bookingapp/models"
)

const (
	QUERY_INSERT_HOTEL = "INSERT INTO hotels (hotel_name, hotel_address, hotel_city, hotel_state, hotel_zip, hotel_country, hotel_price) VALUES($1, $2, $3, $4, $5, $6, $7)RETURNING hotel_id;"
	QUERY_DELETE_HOTEL = "DELETE FROM hotels WHERE hotel_id = $1;"
	QUERY_SELECT_HOTEL = "SELECT * FROM hotels"
	QUERY_UPDATE_HOTEL = "UPDATE hotels SET hotel_name= $1, hotel_address= $2, hotel_city= $3, hotel_state= $4, hotel_zip=$5, hotel_country=$6, hotel_price=$7 WHERE hotel_id=$8;"
)

func AddHotel(c echo.Context) error {
	c.Response().Header().Set("Content-Type", "application/json")
	UriReq := "[" + c.Request().Method + "]" + " [RequestType = Add_Hotel] " + "[" + c.Path() + "] = "

	var hotel models.Hotel
	var hotels []models.Hotel
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
		logging.ErrorLogging(http.StatusInternalServerError, UriReq, err.Error())
		log.Error(err.Error())
		resp := c.JSON(http.StatusInternalServerError, helper.ErrorLog(http.StatusInternalServerError, "Failed while try to execute query", "EXT_REF"))
		return resp
	}
	hotels = append(hotels, hotel)
	dataResponse := models.RespSuccess(hotels, "SUCCESS ADD HOTEL")
	logging.SuccessLogging(UriReq, dataResponse)
	return c.JSON(http.StatusAccepted, dataResponse)
}

func UpdateHotel(c echo.Context) error {
	c.Response().Header().Set("Content-Type", "application/json")
	UriReq := "[" + c.Request().Method + "]" + " [RequestType = Update_Hotel] " + "[" + c.Path() + "] = "

	var hotel models.Hotel
	var hotels []models.Hotel
	body, _ := ioutil.ReadAll(c.Request().Body)
	err := json.Unmarshal(body, &hotel)
	if err != nil {
		log.Error(err.Error())
		resp := c.JSON(http.StatusNotFound, models.RespFailed(http.StatusNotFound, "Error when unmarshal body"))
		return resp
	}
	var count int64
	DB := db.DB()
	res, _ := DB.Exec(QUERY_UPDATE_HOTEL, &hotel.Name, &hotel.Address, &hotel.City, &hotel.State, &hotel.Zip, &hotel.Country, &hotel.Price, &hotel.HotelId)
	count, _ = res.RowsAffected()
	if count < 1 {
		resp := c.JSON(http.StatusNotFound, models.RespFailed(http.StatusNotFound, "HOTEL NOT FOUND"))
		return resp
	}
	hotels = append(hotels, hotel)
	DataResp := models.RespSuccess(hotels, "SUCCESS UPDATE HOTEL with Id : "+strconv.Itoa(hotel.HotelId))
	logging.SuccessLogging(UriReq, DataResp)
	return c.JSON(http.StatusOK, DataResp)
}

func DeleteHotel(c echo.Context) error {
	c.Response().Header().Set("Content-Type", "application/json")
	UriReq := "[" + c.Request().Method + "]" + " [RequestType = Delete_Hotel] " + "[" + c.Path() + "] = "

	id, _ := strconv.Atoi(c.Param("id"))
	var count int64
	DB := db.DB()
	res, _ := DB.Exec(QUERY_DELETE_HOTEL, id)
	count, _ = res.RowsAffected()
	if count < 1 {
		resp := c.JSON(http.StatusNotFound, models.RespFailed(http.StatusNotFound, "HOTEL NOT FOUND"))
		return resp
	}

	dataResponse := models.RespSuccess(nil, "SUCCESS DELETE HOTEL with Id : "+strconv.Itoa(id))
	logging.SuccessLogging(UriReq, dataResponse)
	return c.JSON(http.StatusOK, dataResponse)
}
func GetAllHotel(c echo.Context) error {
	c.Response().Header().Set("Content-Type", "application/json")
	UriReq := "[" + c.Request().Method + "]" + " [RequestType = Get_Hotel] " + "[" + c.Path() + "] = "

	DB := db.DB()
	var hotel models.Hotel

	rows, err := DB.Query(QUERY_SELECT_HOTEL)
	if err != nil {
		log.Error(err.Error())
		resp := c.JSON(http.StatusNotFound, models.RespFailed(http.StatusNotFound, err.Error()))
		return resp
	}
	defer rows.Close()
	var hotels []models.Hotel
	for rows.Next() {
		rows.Scan(&hotel.HotelId, &hotel.Name, &hotel.Address, &hotel.City, &hotel.State, &hotel.Zip, &hotel.Country, &hotel.Price)
		hotels = append(hotels, hotel)
	}
	dataResp := models.RespSuccess(hotels, "SUCCESS GET DATA")
	logging.SuccessLogging(UriReq, dataResp)
	resp := c.JSON(http.StatusOK, dataResp)
	return resp
}
