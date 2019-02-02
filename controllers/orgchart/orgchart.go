package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

// CreateChart will persist a new chart with a root node
// [PUT, POST] /chartId/:chartId
func CreateChart(c echo.Context) error {
	chartID := c.Param("chartId")

	return c.String(http.StatusOK, fmt.Sprintf("CreateChart Chart ID: %v", chartID))
}

// DeleteChart will remove a chart from persistency
// [DELETE] /chartId/:chartId
func DeleteChart(c echo.Context) error {
	chartID := c.Param("chartId")

	return c.String(http.StatusOK, fmt.Sprintf("DeleteChart Chart ID: %v", chartID))
}
