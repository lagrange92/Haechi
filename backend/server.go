package main

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/cors"
)

var e echo.Echo

// CreateServer creates echo server
func CreateServer() {
	e := echo.New()
	c := cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:8080", // haechi web
		},
		AllowCredentials: true,
	})

	e.Use(echo.WrapMiddleware(c.Handler))

	e.GET("/", HandleHome)
	e.GET("/ppltn", HandlePpltn)
}

// StartServer starts echo server
func StartServer() error {
	return e.Start(":1323")
}
