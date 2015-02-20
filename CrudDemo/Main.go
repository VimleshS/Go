package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Employee struct {
	Id   int
	Name string
	Age  int
}

func main() {
	r := gin.Default()
	/*Below Works*/
	v1 := r.Group("/")
	{
		v1.GET("/:EmployeeId", getHandleFunc)
		v1.GET("/", getHandleFunc)
	}
	r.DELETE("/:EmployeeId", deleteHandleFunc)
	r.POST("/", postHandleFunc)
	r.PUT("/", putHandleFunc)

	r.Run(":8080")
}

func getHandleFunc(c *gin.Context) {
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

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	jsonForEmp := make([]Employee, 0)
	for rows.Next() {
		e := new(Employee)
		rows.Scan(&e.Id, &e.Name, &e.Age)
		jsonForEmp = append(jsonForEmp, *e)
	}
	c.JSON(200, jsonForEmp)
}

func deleteHandleFunc(c *gin.Context) {
	id := c.Params.ByName("EmployeeId")
	sqlQuery := "DELETE FROM EMPLOYEE"
	if id != "" {
		sqlQuery = sqlQuery + " WHERE ID = %s"
		sqlQuery = fmt.Sprintf(sqlQuery, id)
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
	res, err := _mainDb.Execute(sqlQuery)
	if err != nil {
		c.String(401, "Execute: "+err.Error())
	}

	rowsAfected, err := res.RowsAffected()
	if rowsAfected > 0 {
		c.JSON(200, gin.H{"Data": fmt.Sprintf("%d rows deleted.", rowsAfected)})
	}
}

func postHandleFunc(c *gin.Context) {
	var json Employee
	c.Bind(&json)

	sqlQuery := "UPDATE EMPLOYEE SET NAME = %s, AGE =%d WHERE ID = %d"
	sqlQuery = fmt.Sprintf(sqlQuery, strconv.Quote(json.Name), json.Age, json.Id)
	_mainDb := new(MainDb)
	err := _mainDb.Init("test")
	if err != nil {
		c.String(401, "Init: "+err.Error())
	} else {
		defer func() {
			_mainDb.Dispose()
		}()
	}
	res, err := _mainDb.Execute(sqlQuery)
	if err != nil {
		c.String(401, "Execute: "+err.Error())
	}
	rowsAfected, err := res.RowsAffected()
	if rowsAfected > 0 {
		c.JSON(200, gin.H{"Data": fmt.Sprintf("%d rows updated.", rowsAfected)})
	}
}

func putHandleFunc(c *gin.Context) {
	var json Employee
	c.Bind(&json)

	_mainDb := new(MainDb)
	err := _mainDb.Init("test")
	if err != nil {
		c.String(401, "Init: "+err.Error())
	} else {
		defer func() {
			_mainDb.Dispose()
		}()
	}
	//fmt.Println("-------------------------------------")
	sqlQuery := "SELECT MAX(ID) + 1 FROM EMPLOYEE"
	rows, err := _mainDb.Read(sqlQuery)
	if err != nil {
		c.String(401, "Execute: "+err.Error())
	}

	rows.Scan(&json.Id)
	//fmt.Println("------------ID -------------------------" + string(json.Id))
	sqlQuery = "INSERT INTO EMPLOYEE( ID, NAME, AGE ) VALUES ( %d, %s, %d)"
	sqlQuery = fmt.Sprintf(sqlQuery, json.Id, strconv.Quote(json.Name), json.Age)
	//fmt.Println(sqlQuery)
	res, err := _mainDb.Execute(sqlQuery)
	if err != nil {
		c.String(401, "Execute: "+err.Error())
	}
	rowsAfected, err := res.RowsAffected()
	if rowsAfected > 0 {
		c.JSON(200, gin.H{"Data": fmt.Sprintf("%d rows inserted.", rowsAfected)})
	}
}
