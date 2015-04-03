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

func main1() {

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
func main() {

	db, err := sql.Open("mysql", "admin:admin#123@/lrs_connect_dev")
	if err != nil {
		fmt.Printf("Error in GetMySQLQuestions %s\n", err.Error())

	}
	defer db.Close()

	rows, err := db.Query("Select survey_question_type,guid,question_text,is_deleted,survey_page_type_id,header_text,body_text from survey_question where survey_id = 1351 order by display_order asc")
	if err != nil {
		fmt.Printf("Error in GetMySQLQuestions %s\n", err.Error())

	}

	for rows.Next() {

		var survey_question_type int
		var guid string
		var question_text string
		var is_deleted int
		var survey_page_type_id int
		//		var header_text string
		//		var body_text string

		var col1, col2 []byte

		err := rows.Scan(&survey_question_type,
			&guid,
			&question_text,
			&is_deleted,
			&survey_page_type_id,
			&col1,
			&col2)

		if err != nil {
			fmt.Printf("ERROR SCANNING %s\n", err.Error())
		}
		fmt.Printf("%-3d %s-20 --> %-70s  %d %d -> %-70s %s\n", survey_question_type, guid, question_text, is_deleted, survey_page_type_id, string(col1), string(col2))

	}

}
