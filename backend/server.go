package main

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/cors"
)

// Server : echo server
type Server struct {
	e *echo.Echo
}

// CreateServer creates echo server
func CreateServer() Server {
	e := echo.New()
	c := cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:8080", // haechi web
		},
		AllowCredentials: true,
	})

	e.Use(echo.WrapMiddleware(c.Handler))

	e.GET("/", HandleHome)
	e.GET("/ppl", HandlePpl)

	return Server{e}
}

// StartServer starts echo server
func StartServer(server Server) error {
	return server.e.Start(":1323")
}
