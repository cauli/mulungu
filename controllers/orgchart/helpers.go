package controllers

import (
	"fmt"
	"net/http"

	"github.com/cauli/mulungu/models"
	"github.com/cauli/mulungu/storage"
	"github.com/cauli/mulungu/tree"
	"github.com/labstack/echo"
)

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
