package main

import (
	"fmt"
	"github.com/ziutek/mymysql/mysql"
	_ "github.com/ziutek/mymysql/native" // Native engine
)

type Employee struct {
	Id   int32
	Name string
	Age  int32
}

func main() {
	db := mysql.New("tcp", "", "127.0.0.1:3306", "testuser", "TestPasswd9", "test")
	db.Connect()

	//rows, res, err := db.Query("SELECT * FROM employee where Id = %d", 2)
	//rows, res, err := db.Query("SELECT * FROM employee")
	//if err != nil {
	//	panic(err)
	//}

	//for _, row := range rows {
	//	fmt.Println(row.Str(0), "  ", row.Str(1))
	//	fmt.Println(row.Str(res.Map("Id")), "  ", row.Str(res.Map("Name")))
	//}

	//data := new(Employee)
	//data.Id = 3
	//data.Name = "AppCreated"
	//data.Age = 55
	fmt.Println("----------------FFFFFFF------------")
	//stmt, err := db.Prepare("Insert into Employee(Id, Name, Age) values (1, ?, ?)")
	stmt, err := db.Prepare("insert into AAAA values (?, ?)")
	fmt.Println(stmt)
	if err != nil {
		panic(err)
	}
	fmt.Println("---------------------------------------------")
	//_, err = stmt.Run(data)
	//if err != nil {
	//	panic(err)
	//}
	fmt.Println("-------------VVVVVVVVVVVVV-------------")
}
