package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lagrange92/Haechi/store"
)

// HandleHome : handle requst to '/'
func HandleHome(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to Haechi Server!")
}

// HandlePpl : handle request to '/ppltn'
func HandlePpl(c echo.Context) error {
	return c.JSON(http.StatusOK, store.PplDistribution)
}
