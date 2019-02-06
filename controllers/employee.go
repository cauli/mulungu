package controllers

import (
	"fmt"
	"net/http"

	"github.com/cauli/mulungu/model"
	"github.com/cauli/mulungu/storage"
	"github.com/cauli/mulungu/tree"
	"github.com/labstack/echo"
)

// GetSubordinates fetches all children nodes of a desired node
// [GET] /chartId/:chartId/employee/:employeeId/subordinates
func GetSubordinates(c echo.Context) error {
	chartID := c.Param("chartId")
	employeeID := c.Param("employeeId")

	chart, err := findChart(chartID)
	if err != nil {
		return err.Handle(c)
	}

	employee, err := findEmployee(employeeID, chart, true)
	if err != nil {
		return err.Handle(c)
	}

	subordinates := (*employee).GetDescendants()

	return c.JSON(http.StatusOK, subordinates)
}

// UpsertEmployee creates or updates an employee information
// it is also possible to update its current leader by sending a `leader` key
// [PUT] /chartId/:chartId/employee/:employeeId
func UpsertEmployee(c echo.Context) error {
	chartID := c.Param("chartId")
	isUpdating := false

	requestEmployee, err := parseRequestEmployee(c)
	if err != nil {
		return err.Handle(c)
	}

	chart, err := findChart(chartID)
	if err != nil {
		return err.Handle(c)
	}

	employee, err := findEmployee(requestEmployee.ID, chart, false)
	if err != nil {
		return err.Handle(c)
	}

	if employee != nil {
		isUpdating = true

		err := updateEmployee(employee, requestEmployee, chart)
		if err != nil {
			return err.Handle(c)
		}
	} else {
		err := createEmployee(requestEmployee, chart)
		if err != nil {
			return err.Handle(c)
		}
	}

	chartJSON, errJSON := chart.ToJSON()
	if errJSON != nil {
		return model.ApiError{errJSON.Error(), http.StatusInternalServerError}.Handle(c)
	}

	storage.Save(resource, chartID, chartJSON)

	if isUpdating {
		return model.ApiResponse{"Employee was successfully updated"}.Handle(c)
	}

	return model.ApiResponse{"Employee was successfully created"}.Handle(c)
}

func updateEmployee(employee *tree.Node, requestEmployee *model.Employee, chart *tree.Tree) *model.ApiError {
	(*employee).Data = tree.MetaData{
		Name:  requestEmployee.Name,
		Title: requestEmployee.Title,
	}

	if (*employee).ParentID != requestEmployee.Leader {
		desiredLeader, err := findEmployee(requestEmployee.Leader, chart, true)
		if err != nil {
			return err
		}

		errMove := chart.MoveNode(employee, desiredLeader)
		if errMove != nil {
			return &model.ApiError{fmt.Sprintf("Could not update employee's new leader.\nDetails:%s", errMove.Error()), http.StatusBadRequest}
		}
	}

	return nil
}

func createEmployee(requestEmployee *model.Employee, chart *tree.Tree) *model.ApiError {
	newEmployee := requestEmployee.CreateNode()

	desiredLeader, err := findEmployee(requestEmployee.Leader, chart, true)
	if err != nil {
		return err
	}

	errAttach := chart.AttachNode(newEmployee, desiredLeader)
	if errAttach != nil {
		return &model.ApiError{fmt.Sprintf("Could not add new employee.\nDetails:%s", errAttach.Error()), http.StatusBadRequest}
	}

	return nil
}
