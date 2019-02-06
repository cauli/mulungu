package controllers

import (
	"fmt"
	"net/http"

	"github.com/cauli/mulungu/storage"
	"github.com/cauli/mulungu/tree"
	"github.com/labstack/echo"
)

const resource = "tree"

// GetChart retrieves a chart and returns it in JSON format
// [GET] /chart/:chartId
func GetChart(c echo.Context) error {
	chartID := c.Param("chartId")

	exists, value := storage.GetById(resource, chartID)
	if !exists {
		return c.String(http.StatusNotFound, fmt.Sprintf("Chart `%v` does not exist", chartID))
	}

	chart, err := tree.FromJSON(value.(string))
	if err != nil {
		return c.String(http.StatusNotFound, fmt.Sprintf("Could not parse value for chart `%v`", chartID))
	}

	return c.JSONPretty(http.StatusOK, chart, "")
}

// CreateChart will persist a new chart with a root node
// [PUT, POST] /chart/:chartId
func CreateChart(c echo.Context) error {
	chartID := c.Param("chartId")

	exists, _ := storage.GetById(resource, chartID)
	if exists {
		return c.String(http.StatusBadRequest, fmt.Sprintf("Chart `%v` already exists", chartID))
	}

	json, err := tree.New(chartID).ToJSON()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	storage.Save(resource, chartID, json)
	return c.String(http.StatusOK, fmt.Sprintf("Chart `%v` was created", chartID))
}

// DeleteChart will remove a chart from persistency
// [DELETE] /chart/:chartId
func DeleteChart(c echo.Context) error {
	chartID := c.Param("chartId")

	deleted := storage.Delete(resource, chartID)

	if deleted {
		return c.String(http.StatusOK, fmt.Sprintf("Chart `%s` was deleted", chartID))
	}

	return c.String(http.StatusNotFound, fmt.Sprintf("Chart `%s` does not exist", chartID))
}
