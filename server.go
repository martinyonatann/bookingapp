package main

import (
	"net/http"
	"time"

	"github.com/martinyonathann/bookingapp/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	v1 := e.Group("/v1")
	groupRouteHotel(v1)
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST},
	}))

	e.POST("/login", controllers.LoginController)
	e.POST("/registration", controllers.Registration)
	e.GET("/profil", controllers.ProfileHandler)

	s := &http.Server{
		Addr:         ":1323",
		ReadTimeout:  20 * time.Minute,
		WriteTimeout: 20 * time.Minute,
	}
	e.Logger.Fatal(e.StartServer(s))
}
func groupRouteHotel(e *echo.Group) {
	hotel := e.Group("/hotel")
	hotel.POST("", controllers.AddHotel)
	hotel.GET("", controllers.GetAllHotel)
	hotel.DELETE("/:id", controllers.DeleteHotel)

}

// func groupRouteBooking (e *echo.Group){
// 	booking := e.Group("/booking")
// 	booking.POST("", controllers)

// }
