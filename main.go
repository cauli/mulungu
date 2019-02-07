package main

import (
	"fmt"
	"sync/atomic"
	"time"

	"github.com/cauli/mulungu/controllers"
	postgres "github.com/cauli/mulungu/storage/driver"
	"github.com/labstack/echo"
)

const color = "\033[1;31m%s\033[0m"
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

	storage := prepareStorage()
	defer storage.Pool.Close()

	fmt.Printf(color, brand)

	e.GET("/chart/:chartId", controllers.GetChart)
	e.PUT("/chart/:chartId", controllers.CreateChart)
	e.DELETE("/chart/:chartId", controllers.DeleteChart)

	e.PUT("/chart/:chartId/employee/:employeeId", controllers.UpsertEmployee)
	e.GET("/chart/:chartId/employee/:employeeId/subordinates", controllers.GetSubordinates)

	e.Logger.Fatal(e.Start(":8080"))
}

func prepareStorage() *postgres.Storage {

	ticker := time.NewTicker(1500 * time.Millisecond)
	done := make(chan *postgres.Storage, 1)

	go func() {
		var retries uint64
		for range ticker.C {
			storage, err := postgres.New("chart")
			if err != nil {
				atomic.AddUint64(&retries, 1)
				fmt.Println(fmt.Sprintf("Error connecting to storage. Retrying... (%v)", atomic.LoadUint64(&retries)))
			} else {
				postgres.SetMainStorage(&storage)
				done <- &storage
			}
		}
	}()

	return <-done
}
