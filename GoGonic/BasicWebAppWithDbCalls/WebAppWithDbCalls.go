package main

import (
	"github.com/gin-gonic/gin"
)

type Employee struct {
	Id   string
	Name string
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {

		_mainDb := new(MainDb)
		err := _mainDb.Init("test")
		if err != nil {
			c.String(401, "Init: "+err.Error())

		} else {
			defer _mainDb.Dispose()
		}

		rows, err := _mainDb.ExecuteQuery("SELECT * FROM employee")
		if err != nil {
			c.String(401, "Execute: "+err.Error())
		}
		for _, row := range rows {
			jsonobj := Employee{row.Str(0), row.Str(1)}
			c.JSON(200, jsonobj)
		}

	})
	r.Run(":8080")
}
