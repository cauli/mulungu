package main

import (
	"./controllers/orgchart"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.PUT("/chartId/:chartId", controllers.CreateChart)
	e.POST("/chartId/:chartId", controllers.CreateChart)

	e.GET("/chartId/:chartId/employee/:employeeId/subordinates", controllers.GetSubordinates)

	e.PUT("/chartId/:chartId/employee/:employeeId/boss/:bossId", controllers.UpdateBoss)
	e.POST("/chartId/:chartId/employee/:employeeId/boss/:bossId", controllers.UpdateBoss)

	e.Logger.Fatal(e.Start(":8080"))
}
