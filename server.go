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

	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST},
	}))

	e.POST("/login", controllers.LoginController)
	e.POST("/registration", controllers.Registration)
	e.GET("/profil", controllers.ProfileHandler)
	e.POST("/hotel",controllers.AddHotel)

	s := &http.Server{
		Addr:         ":1323",
		ReadTimeout:  20 * time.Minute,
		WriteTimeout: 20 * time.Minute,
	}
	e.Logger.Fatal(e.StartServer(s))
}
