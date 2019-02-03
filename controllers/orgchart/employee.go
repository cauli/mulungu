package controllers

import (
	"fmt"
	"net/http"

	"../../storage"
	"../../tree"

	"github.com/labstack/echo"
)

// GetSubordinates fetches all children nodes of a desired node
// [GET] /chartId/:chartId/employee/:employeeId/subordinates
func GetSubordinates(c echo.Context) error {
	chartID := c.Param("chartId")
	employeeID := c.Param("employeeId")

	exists, value := storage.GetById(resource, chartID)
	if !exists {
		return c.String(http.StatusNotFound, fmt.Sprintf("Chart `%v` does not exist", chartID))
	}

	chart, err := tree.FromJSON(value.(string))
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Could not parse value for chart `%v`", chartID))
	}

	employee, err := chart.FindNode(employeeID, nil)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	if employee == nil {
		return c.String(http.StatusNotFound, fmt.Sprintf("Employee ID `%v` not found for chart `%v`", employeeID, chartID))
	}

	subordinates, err := (*employee).GetDescendants()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, subordinates)
}

// UpdateLeader changes the parent node (boss) of a node
// [POST, PUT] /chartId/:chartId/employee/:employeeId/leader/:leaderId
func UpdateLeader(c echo.Context) error {
	chartID := c.Param("chartId")
	employeeID := c.Param("employeeId")
	leaderID := c.Param("leaderId")

	return c.String(http.StatusOK, fmt.Sprintf("UpdateLeader Chart ID: %v\nEmployee ID: %v\nLeader ID: %v", chartID, employeeID, leaderID))
}
