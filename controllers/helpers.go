package controllers

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/cauli/mulungu/model"
	"github.com/cauli/mulungu/storage"
	"github.com/cauli/mulungu/tree"
	"github.com/labstack/echo"
)

func findEmployee(employeeID string, chart *tree.Tree, isFindingRequired bool) (*tree.Node, *model.ApiError) {
	errID := validateID(employeeID)
	if errID != nil {
		return nil, errID
	}

	if chart == nil {
		return nil, &model.ApiError{"Chart was invalid while finding employee", http.StatusBadRequest}
	}

	employee, err := chart.FindNode(employeeID, nil)
	if err != nil {
		return nil, &model.ApiError{err.Error(), http.StatusInternalServerError}
	}

	if isFindingRequired && employee == nil {
		return nil, &model.ApiError{fmt.Sprintf("Employee ID `%v` not found for chart `%v`", employeeID, chart.Id), http.StatusNotFound}
	}

	return employee, nil
}

func findChart(chartID string) (*tree.Tree, *model.ApiError) {
	errID := validateID(chartID)
	if errID != nil {
		return nil, errID
	}

	exists, value := storage.Load(resource, chartID)
	if !exists {
		return nil, &model.ApiError{fmt.Sprintf("Chart `%v` does not exist", chartID), http.StatusNotFound}
	}

	chart, err := tree.FromJSON(value.(string))
	if err != nil {
		return nil, &model.ApiError{fmt.Sprintf("Could not parse value for chart `%v`", chartID), http.StatusInternalServerError}
	}

	return chart, nil
}

func validateID(id string) *model.ApiError {
	match, _ := regexp.MatchString("^[a-zA-Z0-9-]*$", id)

	if !match {
		return &model.ApiError{fmt.Sprintf("Invalid ID: `%s`. Must be alphanumeric, with hyphens allowed.", id), http.StatusBadRequest}
	}

	return nil
}

func parseRequestEmployee(c echo.Context) (*model.Employee, *model.ApiError) {
	employee := new(model.Employee)
	if err := c.Bind(&employee); err != nil {
		return nil, &model.ApiError{"Invalid request body for employee", http.StatusBadRequest}
	}

	employee.ID = c.Param("employeeId")

	err := validateID(employee.ID)
	if err != nil {
		return nil, err
	}

	return employee, nil
}
