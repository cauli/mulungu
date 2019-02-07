package model

import (
	"net/http"

	"github.com/labstack/echo"
)

// ApiError is the model for request errors,
// containing a `status` (http status code) and a message
type ApiError struct {
	Message string `json:"message,omitempty"`
	Code    int    `json:"status,omitempty"`
}

// Handle will serve a JSON for this error, given a request context
func (err ApiError) Handle(c echo.Context) error {
	return c.JSON(err.Code, err)
}

// ApiResponse is the model for request resposne messages
// it contains only a simple message that can be transformed to JSON
type ApiResponse struct {
	Message string `json:"message,omitempty"`
}

// Handle will serve a JSON for this response, given a request context
func (err ApiResponse) Handle(c echo.Context) error {
	return c.JSON(http.StatusOK, err)
}
