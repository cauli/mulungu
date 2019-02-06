package main

import (
	"fmt"

	controllers "github.com/cauli/mulungu/controllers/orgchart"
	"github.com/labstack/echo"
)

const brandColorRed = "\033[1;31m%s\033[0m"
const brand = `
	        .    ..  ..                   
	        ..    ....                    
	        ..     ..                     
	         ...   ..   ...               
	            ....   ..         ..      
	..      .   ...   ..      .....       
	 ..      .....     ..    ..   ..      
	  ..        ...     ......     ....   
	   ..         ..      ..              
	    .......    ..     ..    ..        
	       .....   ..     ..   ..         
	      ...  ..  ..     ......          
	    ...    ..  ..   ......            
	            ..  . ... ..              
	             .. ... ...               
	              .......                 
	                 ..                   
	                 ..                   
	                 ..                   
	               ......                 
	              ........     MULUNGU    `

func main() {
	e := echo.New()
	e.HideBanner = true

	fmt.Printf(brandColorRed, brand)

	e.GET("/chart/:chartId", controllers.GetChart)
	e.PUT("/chart/:chartId", controllers.CreateChart)
	e.DELETE("/chart/:chartId", controllers.DeleteChart)

	e.PUT("/chart/:chartId/employee/:employeeId", controllers.UpsertEmployee)
	e.GET("/chart/:chartId/employee/:employeeId/subordinates", controllers.GetSubordinates)

	e.Logger.Fatal(e.Start(":8080"))
}
