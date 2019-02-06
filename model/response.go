package model

import (
	"net/http"

	"github.com/labstack/echo"
)

type ApiError struct {
	Message string `json:"message,omitempty"`
	Code    int    `json:"status,omitempty"`
}

func (err ApiError) Handle(c echo.Context) error {
	return c.JSON(err.Code, err)
}

type ApiResponse struct {
	Message string `json:"message,omitempty"`
}

func (err ApiResponse) Handle(c echo.Context) error {
	return c.JSON(http.StatusOK, err)
}
