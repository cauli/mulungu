package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"

	"../../storage"
)

const resource = "tree"

// GetChart retrieves a chart
// [GET] /chart/:chartId
func GetChart(c echo.Context) error {
	chartID := c.Param("chartId")

	exists, value := storage.GetById(resource, chartID)

	if !exists {
		return c.String(http.StatusNotFound, fmt.Sprintf("Chart `%v` does not exist", chartID))
	}

	return c.String(http.StatusOK, fmt.Sprintf("Chart `%v` was found: %v", chartID, value))
}

// CreateChart will persist a new chart with a root node
// [PUT, POST] /chart/:chartId
func CreateChart(c echo.Context) error {
	chartID := c.Param("chartId")

	value, err := storage.GetById(resource, chartID)
	if err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprintf("Chart `%v` already exists", chartID))
	}

	if value {
		return c.String(http.StatusBadRequest, fmt.Sprintf("Chart `%v` already exists", chartID))
	}

	storage.Save(resource, chartID, "test")
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
