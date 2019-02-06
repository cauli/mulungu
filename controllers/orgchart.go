package controllers

import (
	"fmt"
	"net/http"

	"github.com/cauli/mulungu/model"
	"github.com/cauli/mulungu/storage"
	"github.com/cauli/mulungu/tree"
	"github.com/labstack/echo"
)

const resource = "tree"

// GetChart retrieves a chart and returns it in JSON format
// [GET] /chart/:chartId
func GetChart(c echo.Context) error {
	chartID := c.Param("chartId")

	errID := validateID(chartID)
	if errID != nil {
		return errID.Handle(c)
	}

	exists, value := storage.GetById(resource, chartID)
	if !exists {
		message := fmt.Sprintf("Chart `%v` does not exist", chartID)
		return model.ApiError{message, http.StatusNotFound}.Handle(c)
	}

	chart, err := tree.FromJSON(value.(string))
	if err != nil {
		message := fmt.Sprintf("Could not parse value for chart `%v`", chartID)
		return model.ApiError{message, http.StatusInternalServerError}.Handle(c)
	}

	return c.JSONPretty(http.StatusOK, chart, "")
}

// CreateChart will persist a new chart with a root node
// [PUT, POST] /chart/:chartId
func CreateChart(c echo.Context) error {
	chartID := c.Param("chartId")

	errID := validateID(chartID)
	if errID != nil {
		return errID.Handle(c)
	}

	exists, _ := storage.GetById(resource, chartID)
	if exists {
		message := fmt.Sprintf("Chart `%v` already exists", chartID)
		return model.ApiError{message, http.StatusBadRequest}.Handle(c)
	}

	json, err := tree.New(chartID).ToJSON()
	if err != nil {
		return model.ApiError{"Could not unmarshall tree", http.StatusInternalServerError}.Handle(c)
	}

	storage.Save(resource, chartID, json)

	return model.ApiResponse{fmt.Sprintf("Chart `%v` was created", chartID)}.Handle(c)
}

// DeleteChart will remove a chart from persistency
// [DELETE] /chart/:chartId
func DeleteChart(c echo.Context) error {
	chartID := c.Param("chartId")

	errID := validateID(chartID)
	if errID != nil {
		return errID.Handle(c)
	}

	deleted := storage.Delete(resource, chartID)

	if deleted {
		return model.ApiResponse{fmt.Sprintf("Chart `%s` was successfully deleted", chartID)}.Handle(c)
	}

	return model.ApiError{fmt.Sprintf("Chart `%s` does not exist", chartID), http.StatusNotFound}.Handle(c)
}
