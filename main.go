package main

import (
	"./controllers/orgchart"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.GET("/chart/:chartId", controllers.GetChart)
	e.PUT("/chart/:chartId", controllers.CreateChart)
	e.POST("/chart/:chartId", controllers.CreateChart)
	e.DELETE("/chart/:chartId", controllers.DeleteChart)

	e.GET("/chart/:chartId/employee/:employeeId/subordinates", controllers.GetSubordinates)

	e.PUT("/chart/:chartId/employee/:employeeId", controllers.CreateEmployee)

	e.PUT("/chart/:chartId/employee/:employeeId/leader/:leaderId", controllers.UpdateLeader)
	e.POST("/chart/:chartId/employee/:employeeId/leader/:leaderId", controllers.UpdateLeader)

	e.Logger.Fatal(e.Start(":8080"))
}
