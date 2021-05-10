package controllers

// func CreateBooking(c echo.Context) error{
// 	c.Response().Header().Set("Content-Type","application/json")
// 	UriReq := "[" + c.Request().Method + "]" + " [RequestType = Add_Booking] " + "[" + c.Path() + "] = "

// 	var booking models.Booking
// 	body, _ := ioutil.ReadAll(c.Request().Body)
// 	err := json.Unmarshal(body, &booking)

// 	if err != nil {
// 		resp := c.JSON(http.StatusInternalServerError, helper.ErrorLog(http.StatusInternalServerError, "Error when unmarshar request", "EXT_REF"))
// 		return resp
// 	}
// 	DB := db.DB();
// 	err = DB.QueryRow()

// }
