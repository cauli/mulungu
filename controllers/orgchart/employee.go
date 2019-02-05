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

// CreateEmployee creates or updates an employee information
// [PUT] /chartId/:chartId/employee/:employeeId
func CreateEmployee(c echo.Context) error {
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
		newEmployee := tree.Node{
			ID: requestEmployee.ID,
			Data: tree.MetaData{
				Name:  requestEmployee.Name,
				Title: requestEmployee.Title,
			},
			ParentID: requestEmployee.Leader,
		}

		desiredLeader, apiErr := findEmployee(requestEmployee.Leader, chart, true)
		if apiErr != nil {
			return c.String(apiErr.Code, apiErr.Message)
		}

		err = chart.AttachNode(&newEmployee, desiredLeader)
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

// UpdateLeader changes the leader of an employee.
// all subordinates from the employee will keep their original leader
// [PUT] /chartId/:chartId/employee/:employeeId/leader/:leaderId
func UpdateLeader(c echo.Context) error {
	chartID := c.Param("chartId")
	employeeID := c.Param("employeeId")
	leaderID := c.Param("leaderId")

	chart, apiErr := findChart(chartID)
	if apiErr != nil {
		return c.String(apiErr.Code, apiErr.Message)
	}

	employee, apiErr := findEmployee(employeeID, chart, true)
	if apiErr != nil {
		return c.String(apiErr.Code, apiErr.Message)
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

func findEmployee(employeeID string, chart *tree.Tree, isFindingRequired bool) (*tree.Node, *models.ApiError) {
	if chart == nil {
		return nil, &models.ApiError{nil, "Chart was invalid while finding employee", http.StatusBadRequest}
	}

	employee, err := chart.FindNode(employeeID, nil)
	if err != nil {
		return nil, &models.ApiError{nil, err.Error(), http.StatusInternalServerError}
	}

	if isFindingRequired && employee == nil {
		return nil, &models.ApiError{nil, fmt.Sprintf("Employee ID `%v` not found for chart `%v`", employeeID, chart.Id), http.StatusNotFound}
	}

	return employee, nil
}

func findChart(chartID string) (*tree.Tree, *models.ApiError) {
	exists, value := storage.GetById(resource, chartID)
	if !exists {
		return nil, &models.ApiError{nil, fmt.Sprintf("Chart `%v` does not exist", chartID), http.StatusNotFound}
	}

	chart, err := tree.FromJSON(value.(string))
	if err != nil {
		return nil, &models.ApiError{nil, fmt.Sprintf("Could not parse value for chart `%v`", chartID), http.StatusInternalServerError}
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
