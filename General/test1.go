package main

import (
	"fmt"
	"github.com/robfig/config"
	"os"
	"path/filepath"
	"time"
)

func main() {
	/*
		test, err := time.Parse("01/02/2006", "01/01/1970")
		if err != nil {
			panic(err)
		}
	*/

	test := time.Date(
		1970, 01, 01, 0, 0, 0, 0, time.UTC)
	fmt.Println(test)
	//	fmt.Println(time.Since(test).UnixNano())

	old := test.UnixNano() / int64(time.Millisecond)
	new := time.Now().UnixNano() / int64(time.Millisecond)
	res := new - old
	fmt.Printf("%T", res)

	myconf, err := config.ReadDefault("app.properties")
	if err != nil {
		panic(err)
	}
	var Conf map[string]string
	Conf = make(map[string]string)
	var keys = []string{"MONGO", "MYSQL"}
	for key := range keys {
		Conf[keys[key]], err = myconf.String("", keys[key])
		if err != nil {
			panic(err)
		}
	}
	fmt.Println()
	fmt.Println(Conf["MONGO"])

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dir)
	f, _ := os.Create("FileName.txt")
	f.WriteString("Hello")
	f.Sync()
	f.Close()

}
