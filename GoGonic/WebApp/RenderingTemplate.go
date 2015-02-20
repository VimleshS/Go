package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	//	"strconv"
)

type Employee struct {
	Id   int
	Name string
	Age  int
}

func main() {
	r := gin.Default()

	r.GET("/test", func(c *gin.Context) {
		//		g.HTML(200, "hello.tmpl", Data{"@dorijastyle"})

		id := c.Params.ByName("EmployeeId")
		//id = "1"
		fmt.Println("Value of Id : " + id)
		if id == "favicon.ico" {
			id = ""
		}

		_mainDb := new(MainDb)
		err := _mainDb.Init("test")
		if err != nil {
			c.String(401, "Init: "+err.Error())
		} else {
			defer func() {
				_mainDb.Dispose()
			}()
		}

		sqlQuery := "SELECT * FROM EMPLOYEE"
		if id != "" {
			sqlQuery = sqlQuery + " WHERE ID = %s"
			sqlQuery = fmt.Sprintf(sqlQuery, id)
		}
		fmt.Println(sqlQuery)

		rows, err := _mainDb.ReadAll(sqlQuery)
		if err != nil {
			c.String(401, "Execute: "+err.Error())
		}

		//c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

		//jsonForEmp := make([]Employee, 0)
		jsonForEmp := []Employee{}
		for rows.Next() {
			e := new(Employee)
			rows.Scan(&e.Id, &e.Name, &e.Age)
			jsonForEmp = append(jsonForEmp, *e)
		}

		fmt.Println("-----------------------------")
		//r.LoadHTMLFiles("layout.tmpl", "Emp.tmpl")
		temp := template.Must(template.ParseFiles("layout.tmpl", "Emp.tmpl"))
		r.SetHTMLTemplate(temp)
		c.HTML(200, "base", jsonForEmp)
	})

	r.Run(":8080")
}
