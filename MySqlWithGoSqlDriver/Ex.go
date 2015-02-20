package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func printOK() {
	fmt.Println("OK")
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {

	/*Both Works*/
	//db, err := sql.Open("mysql", "testuser:TestPasswd9@/test")
	/*Both Works*/
	db, err := sql.Open("mysql", ":@/test")

	checkError(err)
	rows, err := db.Query("select * from Employee")
	checkError(err)
	var id, age int
	var name string
	for rows.Next() {
		rows.Scan(&id, &name, &age)
		fmt.Printf(" Id: %d  Name: %s  Age: %d", id, name, age)
		fmt.Println("")
	}
	_, err = db.Exec("Insert into Employee(Id, Name, Age) values (?, ?, ?)", 102, "Sixy", 23)
	checkError(err)

}
