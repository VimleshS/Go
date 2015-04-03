package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	//	"log"
	"reflect"
	"strings"
)

type Person struct {
	Name  string
	Phone string
}

//func main() {
//	//	session, err := mgo.Dial("server1.example.com,server2.example.com")
//	session, err := mgo.Dial("localhost:27017")
//	if err != nil {
//		panic(err)
//	}
//	defer session.Close()
//	fmt.Println("Connected....")

//	// Optional. Switch the session to a monotonic behavior.
//	session.SetMode(mgo.Monotonic, true)

//	c := session.DB("test").C("people")
//	//	err = c.Insert(&Person{"Ale", "+55 53 8116 9639"},
//	//		&Person{"Cla", "+55 53 8402 8510"})
//	//	if err != nil {
//	//		log.Fatal(err)
//	//	}

//	result := Person{}
//	err = c.Find(bson.M{"name": "Ale"}).One(&result)
//	if err != nil {
//		log.Fatal(err)
//	}

//	//Update the fields..
//	err = c.Update(bson.M{"name": "Ale"}, bson.M{"$set": bson.M{"phone": "+56565656"}})
//	if err != nil {
//		log.Fatal(err)
//	}

//	fmt.Println("Phone:", result.Phone)
//	fmt.Printf("%T\n", reflect.TypeOf(result).String())
//}

func main() {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	fmt.Println("Connected....")

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	type Gendata map[string]interface{}

	//data := []Gendata{}
	data := []MainDoc{}
	c := session.DB("reports").C("surveyResponse")
	//	err = c.Find(bson.M{"document.surveyGuid": "83b802b1-23c3-42de-b542-0455351e9ff3",
	//		"document.dateParts.date": bson.M{"$gte": 20140220, "$lte": 20150220},
	//		"document.questionGuid":   "5ddef910-8ea2-4830-8e58-543ce57ed011"}).All(&data)

	//All+New+Brand+Survey&dateFrom=2014-03-05&dateTo=2015-03-27

	err = c.Find(bson.M{"document.surveyGuid": "83b802b1-23c3-42de-b542-0455351e9ff3",
		"document.dateParts.date": bson.M{"$gte": 20140305, "$lte": 20150327},
	}).All(&data)

	if err != nil {
		fmt.Printf("error %s", err.Error())
	}
	//	for _, v := range data {
	//		if v.ReceivedDocument.SurveyResponseId == "63dcd472-a4c1-451f-bba7-b2260df3eb57" {
	//			fmt.Println(v.ReceivedDocument.Answer)
	//			fmt.Println(v.ReceivedDocument.Question)
	//			fmt.Println(v.ReceivedDocument.Validation)
	//			fmt.Println("---------------------------------")
	//		}
	//	}

	rows := make(map[string]bool)
	for _, v := range data {
		_, exist := rows[v.ReceivedDocument.SurveyResponseId]
		if exist == false {
			rows[v.ReceivedDocument.SurveyResponseId] = false
			fmt.Println(v.ReceivedDocument.SurveyResponseId)
		}
	}
	fmt.Println("---------------------------------")

	//	for _, v := range data {
	//		printValues(v)
	//	}
}

func printValues(data map[string]interface{}) {
	fmt.Println(strings.Repeat("-", 76))
	for i, v := range data {
		if reflect.TypeOf(v).String() == "map[string]interface {}" {
			printValues(v.(map[string]interface{}))
		} else {
			fmt.Printf("| %-20s|  %-50v|  \n", i, v)
		}
	}
	fmt.Println(strings.Repeat("-", 76))
}
