package controllers

import (
	"fmt"
	"net/http"

	"../../models"
	"../../storage"
	"../../tree"

	"github.com/labstack/echo"
)

// GetSubordinates fetches all children nodes of a desired node
// [GET] /chartId/:chartId/employee/:employeeId/subordinates
func GetSubordinates(c echo.Context) error {
	chartID := c.Param("chartId")
	employeeID := c.Param("employeeId")

	chart, err := findChart(c, chartID)
	if err != nil {
		return err
	}

	employee, err := findEmployee(c, employeeID, chart)
	if err != nil {
		return err
	}

	subordinates, err := (*employee).GetDescendants()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, subordinates)
}

// CreateEmployee creates or updates an employee information
// [PUT] /chartId/:chartId/employee/:employeeId
func CreateEmployee(c echo.Context) error {
	chartID := c.Param("chartId")

	requestEmployee, err := parseRequestEmployee(c)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid request body for Employee")
	}

	chart, err := findChart(c, chartID)
	if err != nil {
		return err
	}

	employee, err := findEmployee(c, requestEmployee.ID, chart)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, employee)
}

// UpdateLeader changes the leader of an employee.
// all subordinates from the employee will keep their original leader
// [PUT] /chartId/:chartId/employee/:employeeId/leader/:leaderId
func UpdateLeader(c echo.Context) error {
	chartID := c.Param("chartId")
	employeeID := c.Param("employeeId")
	leaderID := c.Param("leaderId")

	chart, err := findChart(c, chartID)
	if err != nil {
		return err
	}

	employee, err := findEmployee(c, employeeID, chart)
	if err != nil {
		return err
	}

	leader, err := chart.FindNode(leaderID, nil)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	if leader == nil {
		return c.String(http.StatusNotFound, fmt.Sprintf("Leader ID `%v` not found for chart `%v`", employeeID, chartID))
	}

	err = (*chart).MoveNode(employee, leader)
	if err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprintf("An error has occured while trying to update the leader of employee ID '%s'.\nDetails: %s", employeeID, err.Error()))
	}

	return c.String(http.StatusOK, fmt.Sprintf("UpdateLeader Chart ID: %v\nEmployee ID: %v\nLeader ID: %v", chartID, employeeID, leaderID))
}

func findEmployee(c echo.Context, employeeID string, chart *tree.Tree) (*tree.Node, error) {
	if chart == nil {
		return nil, c.String(http.StatusBadRequest, "Chart was invalid while finding employee")
	}

	employee, err := chart.FindNode(employeeID, nil)

	if err != nil {
		return nil, c.String(http.StatusInternalServerError, err.Error())
	}

	if employee == nil {
		return nil, c.String(http.StatusNotFound, fmt.Sprintf("Employee ID `%v` not found for chart `%v`", employeeID, chart.Id))
	}

	return employee, nil
}

func findChart(c echo.Context, chartID string) (*tree.Tree, error) {
	exists, value := storage.GetById(resource, chartID)
	if !exists {
		return nil, c.String(http.StatusNotFound, fmt.Sprintf("Chart `%v` does not exist", chartID))
	}

	chart, err := tree.FromJSON(value.(string))
	if err != nil {
		return nil, c.String(http.StatusInternalServerError, fmt.Sprintf("Could not parse value for chart `%v`", chartID))
	}

	return chart, nil
}

func parseRequestEmployee(c echo.Context) (*models.Employee, error) {
	employee := new(models.Employee)
	if err := c.Bind(&employee); err != nil {
		return &models.Employee{}, err
	}

	employee.ID = c.Param("employeeId")

	return employee, nil
}
