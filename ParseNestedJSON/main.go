package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {
	encoded := "{" +
		`  "name": "Cain",` +
		`  "parents": {` +
		`    "mother": "Eve",` +
		`    "father": "Adam",` +
		`    "add" : {` +
		`             "City" : "Mumbai", ` +
		`             "State" : "Maharashtra" ` +
		`            }` +
		`  }` +
		`}`

	// Decode the json object
	var j map[string]interface{}
	err := json.Unmarshal([]byte(encoded), &j)
	if err != nil {
		panic(err)
	}
	PrintValues(j)
	//	// pull out the parents object
	//	parents := j["parents"].(map[string]interface{})

	//	// Print out mother and father
	//	fmt.Printf("Mother: %s\n", parents["mother"].(string))
	//	fmt.Printf("Father: %s\n", parents["father"].(string))
}

type GenericJSON map[string]interface{}

func PrintValues(data map[string]interface{}) {
	for i, v := range data {
		if reflect.TypeOf(v).String() == "map[string]interface {}" {
			PrintValues(v.(map[string]interface{}))
		} else {
			fmt.Printf(" %-20s |  %-20v|  \n", i, v)
		}
	}
}
