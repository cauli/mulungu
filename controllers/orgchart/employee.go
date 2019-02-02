package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

// GetSubordinates fetches all children nodes of a desired node
// [GET] /chartId/:chartId/employee/:employeeId/subordinates
func GetSubordinates(c echo.Context) error {
	chartID := c.Param("chartId")
	employeeID := c.Param("employeeId")

	return c.String(http.StatusOK, fmt.Sprintf("GetSubordinates Chart ID: %v\n Employee ID: %v", chartID, employeeID))
}

// UpdateBoss changes the parent node (boss) of a node
// [POST, PUT] /chartId/:chartId/employee/:employeeId/boss/:bossId
func UpdateBoss(c echo.Context) error {
	chartID := c.Param("chartId")
	employeeID := c.Param("employeeId")
	bossID := c.Param("bossId")

	return c.String(http.StatusOK, fmt.Sprintf("UpdateBoss Chart ID: %v\nEmployee ID: %v\nBoss ID: %v", chartID, employeeID, bossID))
}
