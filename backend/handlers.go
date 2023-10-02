package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lagrange92/Haechi/chat"
	"github.com/lagrange92/Haechi/model"
	"github.com/lagrange92/Haechi/store"
)

// HandleHome : handle requst to '/'
func HandleHome(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to Haechi Server!")
}

// HandlePpl : handle request to '/ppl'
func HandlePpl(c echo.Context) error {
	return c.JSON(http.StatusOK, store.PplDistribution)
}

// HandleCozy : handle request to '/cozy'
func HandleCozy(c echo.Context) error {
	return c.JSON(http.StatusOK, store.CozyPlaces)
}

// HandleChat : handle request to '/chat'
func HandleChat(c echo.Context) error {
	body := new(model.ChatPromptData)
	if err := c.Bind(body); err != nil {
		return err
	}

	if body.Prompt == "" {
		return c.JSON(http.StatusBadRequest, "prompt field is empty.")
	}

	// for test
	// return c.JSON(http.StatusOK, "제 이름은 AIChat이에요. 저는 인공지능 챗봇이에요. 어떻게 도와드릴까요?")

	resp, err := chat.Chat(body.Prompt)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}
