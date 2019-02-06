package controllers

import (
	"fmt"
	"net/http"

	"github.com/cauli/mulungu/storage"
	"github.com/cauli/mulungu/tree"
	"github.com/labstack/echo"
)

// GetSubordinates fetches all children nodes of a desired node
// [GET] /chartId/:chartId/employee/:employeeId/subordinates
func GetSubordinates(c echo.Context) error {
	chartID := c.Param("chartId")
	employeeID := c.Param("employeeId")

	chart, apiErr := findChart(chartID)
	if apiErr != nil {
		return c.String(apiErr.Code, apiErr.Message)
	}

	employee, apiErr := findEmployee(employeeID, chart, true)
	if apiErr != nil {
		return c.String(apiErr.Code, apiErr.Message)
	}

	subordinates, err := (*employee).GetDescendants()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, subordinates)
}

// UpsertEmployee creates or updates an employee information
// it is also possible to update its current leader by sending a `leader` key`
// [PUT] /chartId/:chartId/employee/:employeeId
func UpsertEmployee(c echo.Context) error {
	chartID := c.Param("chartId")
	isUpdating := false

	requestEmployee, err := parseRequestEmployee(c)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid request body for Employee")
	}

	chart, apiErr := findChart(chartID)
	if apiErr != nil {
		return c.String(apiErr.Code, apiErr.Message)
	}

	employee, apiErr := findEmployee(requestEmployee.ID, chart, false)
	if apiErr != nil {
		return c.String(apiErr.Code, apiErr.Message)
	}

	if employee != nil {
		isUpdating = true

		(*employee).Data = tree.MetaData{
			Name:  requestEmployee.Name,
			Title: requestEmployee.Title,
		}

		if (*employee).ParentID != requestEmployee.Leader {
			desiredLeader, apiErr := findEmployee(requestEmployee.Leader, chart, true)
			if apiErr != nil {
				return c.String(apiErr.Code, apiErr.Message)
			}

			err = chart.MoveNode(employee, desiredLeader)
			if err != nil {
				return c.String(http.StatusBadRequest, fmt.Sprintf("Could not update employee's new leader.\nDetails:%s", err.Error()))
			}
		}
	} else {
		newEmployee := requestEmployee.CreateNode()

		desiredLeader, apiErr := findEmployee(requestEmployee.Leader, chart, true)
		if apiErr != nil {
			return c.String(apiErr.Code, apiErr.Message)
		}

		err = chart.AttachNode(newEmployee, desiredLeader)
		if err != nil {
			return c.String(http.StatusBadRequest, fmt.Sprintf("Could not add new employee to the chart.\nDetails:%s", err.Error()))
		}
	}

	chartJSON, err := chart.ToJSON()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	storage.Save(resource, chartID, chartJSON)

	if isUpdating {
		return c.JSON(http.StatusOK, "Employee was successfully updated")
	}

	return c.JSON(http.StatusOK, "Employee was successfully created")
}
