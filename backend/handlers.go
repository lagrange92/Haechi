package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lagrange92/Haechi/singleton"
)

// HandleHome : handle requst to '/'
func HandleHome(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to Haechi Server!")
}

// HandlePpltn : handle request to '/ppltn'
func HandlePpltn(c echo.Context) error {
	return c.JSON(http.StatusOK, singleton.LatestPpl)
}
